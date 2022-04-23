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
)

from stacks.nuke_handler_cross_role import NukeHandlerCrossRole

from variables import Enviroments


class GraphQLEndpoint(Construct):
    def __init__(
        self,
        scope: Construct,
        construct_id: str,
        multi_cloud_table: dynamodb.Table,
        roles: List[NukeHandlerCrossRole],
        queue: sqs.Queue,
        enviroment: Enviroments,
    ) -> None:
        super().__init__(scope, construct_id)

        s3Bucket = s3.Bucket.from_bucket_name(
            self, id="cdktoolkit-stagingbucket", bucket_name="cdk-hnb659fds-assets-063661473261-eu-central-1"
        )

        sha = "main"
        try:
            with open("./api/sha", "r") as f:
                sha = f.read().rstrip()
        except:
            print("File Not Found, use main")

        environment = {}
        if roles != None:
            for role in roles:
                environment[str(role.id).replace("-", "")] = role.role_name

        func = lambda_.Function(
            self,
            "lambda",
            architecture=lambda_.Architecture.X86_64,
            runtime=lambda_.Runtime.GO_1_X,
            code=lambda_.Code.from_bucket(s3Bucket, f"lambda/{sha}.zip"),
            handler="main",
            timeout=Duration.seconds(30),
            memory_size=128,
            environment={
                **environment,
                "multi_cloud_table": multi_cloud_table.table_name,
                "gitlab_azure_pipeline_webhook": os.getenv("GITLAB_AZURE_PIPELINE_WEBHOOK", "NA"),
                "sqs_sso_assignment": queue.queue_name,
            },
        )

        queue.grant_send_messages(func)

        multi_cloud_table.grant_read_write_data(func)

        api = apigw.LambdaRestApi(
            self,
            "GraphqlApi",
            handler=func,
            proxy=False,
            default_cors_preflight_options=apigw.CorsOptions(
                allow_origins=apigw.Cors.ALL_ORIGINS, allow_methods=apigw.Cors.ALL_METHODS
            ),
        )

        ssm_sandbox_domain_uri = ssm.StringParameter(
            self,
            "SandboxDomainUri",
            description="Name of the API URI",
            parameter_name=f"/{enviroment.value}/sandbox/apiuri",
            string_value=api.url,
        )

        items = api.root.add_resource("graphql")
        items.add_method("POST")  # POST /sandbox

        self.func = func
        self.ssm_sandbox_domain_uri = ssm_sandbox_domain_uri
