#!/usr/bin/env python3

# For consistency with TypeScript code, `cdk` is the preferred import name for
# the CDK's core module.  The following line also imports it as `core` for use
# with examples from the CDK Developer's Guide, which are in the process of
# being updated to use `cdk`.  You may delete this import if you don't need it.
from typing import List
from aws_cdk import (
    App,
    Environment,
    Stage,
    Tags,
)

from stacks.cloud_sandboxes import CloudSandboxes
from stacks.nuke_handler_cross_role import NukeHandlerCrossRole
from stacks.sso_handler_cross_role import SSOHandlerCrossRole
from stacks.monitoring import Monitoring
from stacks.cicd_preperation import CICDPreperation

from variables import sandboxes, root_account, region, sso_account, Enviroments

env_EU = Environment(account=root_account, region=region)

app = App()


class SandboxStage(Stage):
    def __init__(self, scope, id: str, enviroment: Enviroments, *, env=None, outdir=None):
        super().__init__(scope, id, env=env, outdir=outdir)

        Tags.of(self).add("needUntil", "2099-01-01T00:00:00.000Z")
        Tags.of(self).add("creator", "maximilian.haensel@pexon-consulting.de")
        Tags.of(self).add("app", "aws-sandbox-handler")
        Tags.of(self).add("stage", id)

        if enviroment == Enviroments.prod:
            CICDPreperation(self, "CICDPreperation")

        # role_array = NukeHandlerCrossRole.loop(scope=self, sandboxes=sandboxes, enviroment=enviroment)
        nuke_roles = None
        if enviroment == Enviroments.prod:
            nuke_roles: List[NukeHandlerCrossRole] = []
            for account in sandboxes:
                # x = NukeHandlerCrossRole.loop(self, "asda")
                role = NukeHandlerCrossRole(
                    self,
                    account["name"],
                    env={
                        "account": account["id"],
                        "region": region,
                    },
                )
                nuke_roles.append(role)

        sso_role = None
        if enviroment == Enviroments.prod:
            sso_role = SSOHandlerCrossRole(
                self,
                "SSOCrossRole",
                env={
                    "account": sso_account,
                    "region": region,
                },
            )

        sandbox = CloudSandboxes(self, "Sandbox", roles=nuke_roles, sso_role=sso_role, enviroment=enviroment)
        if nuke_roles:
            for role in nuke_roles:
                sandbox.add_dependency(role)

        monitoring = Monitoring(self, "Monitoring", functions=sandbox.functions)
        monitoring.add_dependency(sandbox)


SandboxStage(app, Enviroments.prod.value, env=env_EU, enviroment=Enviroments.prod)
SandboxStage(app, Enviroments.test.value, env=env_EU, enviroment=Enviroments.test)

app.synth()
