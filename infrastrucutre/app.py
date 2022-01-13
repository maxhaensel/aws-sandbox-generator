#!/usr/bin/env python3

# For consistency with TypeScript code, `cdk` is the preferred import name for
# the CDK's core module.  The following line also imports it as `core` for use
# with examples from the CDK Developer's Guide, which are in the process of
# being updated to use `cdk`.  You may delete this import if you don't need it.
from aws_cdk import core
from aws_cdk import (
    aws_iam as iam,
    aws_events as events,
    aws_lambda as lambda_,
    aws_apigateway as apigw,
    aws_dynamodb as dynamodb,
    aws_events_targets as targets,
    aws_ssm as ssm,
)
from aws_cdk.aws_lambda_event_sources import DynamoEventSource

from hosting import AWSSandBoxHosting
from lambda_graphql_endpoint import AWSLambdaGoGraphql

env_EU = core.Environment(account="172920935848", region="eu-central-1")

app = core.App()


class AWSSandboxHandler(core.Stack):
    def __init__(self, scope: core.Construct, id: str, **kwargs) -> None:
        super().__init__(scope, id, **kwargs)

        cicd_user_iac = iam.User(self, "cicd-user-iac")

        cicd_user_iac.add_managed_policy(iam.ManagedPolicy.from_aws_managed_policy_name("AdministratorAccess"))

        table = dynamodb.Table(
            self,
            "Table",
            partition_key=dynamodb.Attribute(name="account_id", type=dynamodb.AttributeType.STRING),
            stream=dynamodb.StreamViewType.NEW_AND_OLD_IMAGES,
            billing_mode=dynamodb.BillingMode.PAY_PER_REQUEST,
        )

        endpoint_lambda = lambda_.Function(
            self,
            "Endpoint_handler",
            architecture=lambda_.Architecture.ARM_64,
            runtime=lambda_.Runtime.PYTHON_3_9,
            code=lambda_.Code.from_asset("lambda/endpoint"),
            handler="handler.handler",
            environment={"dynamodb_table": table.table_name, "duration_of_lease_in_days": "2"},
            timeout=core.Duration.seconds(30),
            memory_size=128,
        )

        table.grant_read_write_data(endpoint_lambda)

        api = apigw.LambdaRestApi(
            self,
            "Endpoint",
            handler=endpoint_lambda,
            proxy=False,
            default_cors_preflight_options=apigw.CorsOptions(
                allow_origins=apigw.Cors.ALL_ORIGINS, allow_methods=apigw.Cors.ALL_METHODS
            ),
        )

        items = api.root.add_resource("sandbox")
        items.add_method("POST")  # POST /sandbox

        # Grant Access to SSO, Remove Access to SSO, Nuke Account
        SSO_Nuke_handler_lambda = lambda_.Function(
            self,
            "SSO_handler_lambda",
            architecture=lambda_.Architecture.ARM_64,
            runtime=lambda_.Runtime.PYTHON_3_9,
            code=lambda_.Code.from_asset("lambda/sso_assaigment"),
            handler="handler.handler",
            timeout=core.Duration.seconds(900),
            memory_size=1024,
            environment={"dry_run": str(True)},
        )

        SSO_Nuke_handler_lambda.role.add_to_principal_policy(
            iam.PolicyStatement(actions=["identitystore:ListUsers", "sts:*"], resources=["*"])
        )
        SSO_Nuke_handler_lambda.role.add_to_principal_policy(
            iam.PolicyStatement(
                actions=["sso:CreateAccountAssignment", "sso:DeleteAccountAssignment"],
                resources=[
                    "arn:aws:sso:::instance/*",
                    "arn:aws:sso:::permissionSet/*/*",
                    "arn:aws:sso:::account/*",
                ],
            )
        )

        SSO_Nuke_handler_lambda.add_event_source(
            DynamoEventSource(
                table,
                starting_position=lambda_.StartingPosition.TRIM_HORIZON,
                batch_size=1,
                bisect_batch_on_error=True,
                retry_attempts=10,
            )
        )

        # clear_sso_and_nuke_account_lambda
        sheduler_sandbox_access_lambda = lambda_.Function(
            self,
            "clear_sso",
            architecture=lambda_.Architecture.ARM_64,
            runtime=lambda_.Runtime.PYTHON_3_9,
            code=lambda_.Code.from_asset("lambda/sheduler"),
            handler="handler.handler",
            environment={"dynamodb_table": table.table_name},
            timeout=core.Duration.seconds(60),
            memory_size=128,
        )

        table.grant_read_write_data(sheduler_sandbox_access_lambda)

        # Run every day at 6PM UTC
        # See https://docs.aws.amazon.com/lambda/latest/dg/tutorial-scheduled-events-schedule-expressions.html
        rule = events.Rule(
            self,
            "Rule",
            schedule=events.Schedule.cron(minute="0", hour="0", month="*", week_day="*", year="*"),
        )
        rule.add_target(targets.LambdaFunction(sheduler_sandbox_access_lambda))

        """
        create Graphql-Endpoint
        """
        lambda_go_graphql = AWSLambdaGoGraphql(self, "graph-ql-endpoint", table)

        """
        create Web-App-Hosting 
        """
        hosting = AWSSandBoxHosting(self, "Hosting", ssm_sandbox_domain_uri=lambda_go_graphql.ssm_sandbox_domain_uri)


AWSSandboxHandler(app, "AWSSandbox", env=env_EU)

core.Tags.of(app).add("needUntil", "2099-01-01T00:00:00.000Z")
core.Tags.of(app).add("creator", "maximilian.haensel@pexon-consulting.de")
core.Tags.of(app).add("app", "aws-sandbox-handler")

app.synth()
