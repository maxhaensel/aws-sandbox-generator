import os
from aws_cdk import (
    core as cdk,
    aws_s3 as s3,
    aws_ssm as ssm,
    aws_lambda as lambda_,
    aws_iam as iam,
    aws_cloudfront as cloudfront,
    aws_cloudfront_origins as origins,
    aws_apigateway as apigw,
    aws_dynamodb as dynamodb,
)
from aws_cdk.aws_ec2 import Volume


class AWSLambdaGoGraphql(cdk.Construct):
    def __init__(
        self, scope: cdk.Construct, construct_id: str, table: dynamodb.Table, multi_cloud_table: dynamodb.Table
    ) -> None:
        super().__init__(scope, construct_id)

        s3Bucket = s3.Bucket.from_bucket_name(
            self, id="cdktoolkit-stagingbucket", bucket_name="cdktoolkit-stagingbucket-1rbmmxnlvi129"
        )

        func = lambda_.Function(
            self,
            "graphql_go_lambda",
            architecture=lambda_.Architecture.X86_64,
            runtime=lambda_.Runtime.GO_1_X,
            code=lambda_.Code.from_bucket(s3Bucket, "lambda/main.zip"),
            handler="main",
            timeout=cdk.Duration.seconds(30),
            memory_size=128,
            environment={
                "dynamodb_table": table.table_name,
                "multi_cloud_table": multi_cloud_table.table_name,
                "gitlab_azure_pipeline_webhook": os.getenv("GITLAB_AZURE_PIPELINE_WEBHOOK", ""),
            },
        )

        table.grant_read_write_data(func)
        multi_cloud_table.grant_read_write_data(func)

        api = apigw.LambdaRestApi(
            self,
            "graphql_go_api",
            handler=func,
            proxy=False,
            default_cors_preflight_options=apigw.CorsOptions(
                allow_origins=apigw.Cors.ALL_ORIGINS, allow_methods=apigw.Cors.ALL_METHODS
            ),
        )

        ssm_sandbox_domain_uri = ssm.StringParameter(
            self,
            "sandbox-domain-uri",
            description="Name of the API URI",
            parameter_name="sandboxDomainUri",
            string_value=api.url,
        )

        items = api.root.add_resource("graphql")
        items.add_method("POST")  # POST /sandbox

        self.ssm_sandbox_domain_uri = ssm_sandbox_domain_uri
