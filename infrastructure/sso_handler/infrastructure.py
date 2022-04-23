from os import environ
from constructs import Construct
from aws_cdk import (
    Duration,
    aws_sqs as sqs,
    aws_iam as iam,
    aws_events as events,
    aws_lambda as lambda_,
    aws_dynamodb as dynamodb,
    aws_events_targets as targets,
)

from aws_cdk.aws_lambda_event_sources import SqsEventSource

from variables import sandboxes, root_account, region

from stacks.sso_handler_cross_role import SSOHandlerCrossRole

timeout = Duration.seconds(900)


class SSOHandler(Construct):
    def __init__(
        self,
        scope: Construct,
        construct_id: str,
        multi_cloud_table=dynamodb.Table,
        sso_role: SSOHandlerCrossRole = None,
        **kwargs
    ) -> None:
        super().__init__(scope, construct_id, **kwargs)

        queue = sqs.Queue(self, "Queue", visibility_timeout=timeout)

        environment = {"table": multi_cloud_table.table_name, "dry_run": str(True)}
        if sso_role != None:
            environment["role_name"] = sso_role.role_name

        # Grant Access to SSO, Remove Access to SSO, Nuke Account
        function = lambda_.Function(
            self,
            "Lambda",
            architecture=lambda_.Architecture.ARM_64,
            runtime=lambda_.Runtime.PYTHON_3_9,
            code=lambda_.Code.from_asset("lambda/sso_assaigment"),
            handler="handler.handler",
            timeout=timeout,
            memory_size=128,
            environment=environment,
        )

        function.add_event_source(SqsEventSource(queue, batch_size=1))

        function.role.add_to_principal_policy(iam.PolicyStatement(actions=["sts:*"], resources=["*"]))

        multi_cloud_table.grant_write_data(function)

        self.func = function
        self.queue = queue


class SandboxGarbageCollector(Construct):
    def __init__(
        self, scope: Construct, construct_id: str, queue: sqs.Queue, sso_role: SSOHandlerCrossRole = None, **kwargs
    ) -> None:
        super().__init__(scope, construct_id, **kwargs)

        function = lambda_.Function(
            self,
            "Lambda",
            architecture=lambda_.Architecture.ARM_64,
            runtime=lambda_.Runtime.PYTHON_3_9,
            code=lambda_.Code.from_asset("lambda/sheduler"),
            handler="handler.handler",
            environment={
                "sqs_sso_assignment": queue.queue_name,
            },
            timeout=Duration.seconds(60),
            memory_size=128,
        )

        queue.grant_send_messages(function)

        # Run every day at 6PM UTC
        # See https://docs.aws.amazon.com/lambda/latest/dg/tutorial-scheduled-events-schedule-expressions.html
        rule = events.Rule(
            self,
            "Rule",
            schedule=events.Schedule.cron(minute="0", hour="0", month="*", week_day="*", year="*"),
        )
        rule.add_target(targets.LambdaFunction(function))

        self.func = function
