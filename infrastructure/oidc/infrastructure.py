import os
from typing import List
from constructs import Construct
from aws_cdk import (
    Duration,
    aws_s3 as s3,
    aws_ssm as ssm,
    aws_lambda as lambda_,
    aws_apigateway as apigw,
    aws_dynamodb as dynamodb,
    aws_sqs as sqs,
    aws_iam as iam,
)

from stacks.nuke_handler_cross_role import NukeHandlerCrossRole

from variables import Enviroments


class GitHubOIDC(Construct):
    def __init__(
        self,
        scope: Construct,
        construct_id: str,
    ) -> None:
        super().__init__(scope, construct_id)

        provider = iam.OpenIdConnectProvider(
            self,
            "GitHubOIDCProvider",
            url="https://token.actions.githubusercontent.com",
            client_ids=["sts.amazonaws.com"],
        )

        role = iam.Role(
            self,
            "OIDCRole",
            assumed_by=iam.FederatedPrincipal(
                federated=provider.open_id_connect_provider_arn,
                assume_role_action="sts:AssumeRoleWithWebIdentity",
                conditions={"StringLike": {"token.actions.githubusercontent.com:sub": "repo:maxhaensel/awsoicd:*"}},
            ),
        )

        role.add_managed_policy(iam.ManagedPolicy.from_aws_managed_policy_name("AdministratorAccess"))
