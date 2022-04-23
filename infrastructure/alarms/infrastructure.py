import os
import typing
from constructs import Construct
from aws_cdk import (
    Duration,
    aws_lambda as lambda_,
    aws_sns as sns,
    aws_cloudwatch as cloudwatch,
    aws_cloudwatch_actions as cw_actions,
    aws_sns_subscriptions as sns_subscriptions,
)


class AWSSlackAlarms(Construct):
    def __init__(
        self,
        scope: Construct,
        construct_id: str,
        functions: typing.List[lambda_.Function],
    ) -> None:
        super().__init__(scope, construct_id)

        alarms: typing.List[cloudwatch.Alarm] = []

        for i, function in enumerate(functions):
            name = function.node.id
            alarm = cloudwatch.Alarm(
                self,
                id=f"{name}Errors{i}",
                alarm_name=f"Error{function.function_name}",
                alarm_description="This alarm occurs when the lambda throw an errors",
                comparison_operator=cloudwatch.ComparisonOperator.GREATER_THAN_THRESHOLD,
                threshold=1,
                evaluation_periods=1,
                metric=function.metric_errors(),
            )
            alarms.append(alarm)

        slack_webhook = lambda_.Function(
            self,
            "slack_webhook",
            architecture=lambda_.Architecture.ARM_64,
            runtime=lambda_.Runtime.PYTHON_3_9,
            code=lambda_.Code.from_asset("lambda/slack"),
            handler="handler.handler",
            timeout=Duration.seconds(60),
            memory_size=128,
            environment={
                "SLACK_WEBHOOK": os.getenv("SLACK_WEBHOOK", ""),
            },
        )

        topic = sns.Topic(self, "Topic", display_name="Customer subscription topic")
        topic.add_subscription(sns_subscriptions.LambdaSubscription(slack_webhook))

        for alarm in alarms:
            alarm.add_alarm_action(cw_actions.SnsAction(topic))
