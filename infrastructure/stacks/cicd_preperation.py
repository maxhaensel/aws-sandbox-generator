from constructs import Construct
from typing import List
from aws_cdk import Stack, aws_iam as iam, aws_lambda as _lambda

from oidc.infrastructure import GitHubOIDC


class CICDPreperation(Stack):
    def __init__(
        self,
        scope: Construct,
        id: str,
        **kwargs,
    ) -> None:
        super().__init__(scope, id, **kwargs)

        """
        create GitHubOIDC
        """
        GitHubOIDC(self, "GitHubOIDC")
