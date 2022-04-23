from typing import List
from constructs import Construct
from aws_cdk import (
    Stack,
    aws_iam as iam,
)

from variables import sandboxes, root_account, region, Enviroments

from database.infrastructure import AWSTable
from hosting.infrastructure import AWSSandBoxHosting
from api.infrastructure import GraphQLEndpoint

from sso_handler.infrastructure import SSOHandler, SandboxGarbageCollector

from stacks.nuke_handler_cross_role import NukeHandlerCrossRole
from stacks.sso_handler_cross_role import SSOHandlerCrossRole


class CloudSandboxes(Stack):
    def __init__(
        self,
        scope: Construct,
        id: str,
        roles: List[NukeHandlerCrossRole],
        enviroment: Enviroments,
        sso_role: SSOHandlerCrossRole = None,
        **kwargs,
    ) -> None:
        super().__init__(scope, id, **kwargs)

        # cicd_user_iac = iam.User(self, "cicd-user-iac")

        # cicd_user_iac.add_managed_policy(iam.ManagedPolicy.from_aws_managed_policy_name("AdministratorAccess"))

        """
        create Graphql-Endpoint
        """
        multi_cloud_table = AWSTable(self, "MultiCloudTable")

        """
        SSO Lambda Handler
        """
        sso_handler = SSOHandler(
            self,
            "SSOHandler",
            sso_role=sso_role,
            multi_cloud_table=multi_cloud_table.table,
        )

        """
        SandboxGarbageCollector
        """
        sandbox_garbage_collector = SandboxGarbageCollector(self, "SandboxGarbageCollector", queue=sso_handler.queue)

        """
        create Graphql-Endpoint
        """
        lambda_go_graphql = GraphQLEndpoint(
            self,
            "GraphQLEndpoint",
            multi_cloud_table=multi_cloud_table.table,
            roles=roles,
            queue=sso_handler.queue,
            enviroment=enviroment,
        )
        """
        create Web-App-Hosting 
        """
        AWSSandBoxHosting(
            self,
            "Hosting",
            ssm_sandbox_domain_uri=lambda_go_graphql.ssm_sandbox_domain_uri,
            enviroment=enviroment,
        )

        self.functions = [lambda_go_graphql.func, sso_handler.func, sandbox_garbage_collector.func]
