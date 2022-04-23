from typing import List
from constructs import Construct
from aws_cdk import (
    Stack,
    PhysicalName,
    aws_iam as iam,
)
from variables import root_account, Enviroments, Sandbox, region


class NukeHandlerCrossRole(Stack):
    def __init__(self, scope: Construct, id: str, **kwargs) -> None:
        super().__init__(scope, id, **kwargs)
        role = iam.Role(
            self,
            f"nuke-sandbox-role-{id}",
            role_name=PhysicalName.GENERATE_IF_NEEDED,
            assumed_by=iam.AccountPrincipal(root_account),
        )
        role.add_managed_policy(iam.ManagedPolicy.from_aws_managed_policy_name("AdministratorAccess"))
        self.role_name = role.role_name
        self.id = id

    # @staticmethod
    # def loop(cls, scope: Construct, sandboxes: List[Sandbox], enviroment: Enviroments, **kwargs):
    #     roles = []
    #     if enviroment == Enviroments.prod:
    #         for sandbox in sandboxes:
    #             role = cls(
    #                 scope,
    #                 sandbox["name"],
    #                 env={
    #                     "account": sandbox["id"],
    #                     "region": region,
    #                 },
    #             )
    #             roles.append(role)
    #     return roles if len(roles) != 0 else None
