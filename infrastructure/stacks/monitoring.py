from constructs import Construct
from typing import List
from aws_cdk import Stack, aws_iam as iam, aws_lambda as _lambda

from alarms.infrastructure import AWSSlackAlarms


class Monitoring(Stack):
    def __init__(
        self,
        scope: Construct,
        id: str,
        functions: List[_lambda.Function],
        **kwargs,
    ) -> None:
        super().__init__(scope, id, **kwargs)

        """
        create Slack Alarms 
        """
        AWSSlackAlarms(self, "SlackAlarms", functions)
