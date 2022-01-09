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


class AWSLambdaGoGraphql(cdk.Construct):
    def __init__(self, scope: cdk.Construct, construct_id: str) -> None:
        super().__init__(scope, construct_id)

        table = dynamodb.Table(
            self,
            "graphql_endpoint_table",
            partition_key=dynamodb.Attribute(name="account_id", type=dynamodb.AttributeType.STRING),
            billing_mode=dynamodb.BillingMode.PAY_PER_REQUEST,
            table_name="sandbox-account-status",
            removalPolicy=cdk.RemovalPolicy.DESTROY,
        )

        endpoint_lambda = lambda_.Function(
            self,
            "graphql_go_lambda",
            architecture=lambda_.Architecture.X86_64,
            runtime=lambda_.Runtime.GO_1_X,
            code=lambda_.Code.from_bucket("cdktoolkit-stagingbucket-1rbmmxnlvi129", "lambda/main.zip"),
            handler="handler.handler",
            timeout=cdk.Duration.seconds(30),
            memory_size=128,
            environment={"dynamodb_table": table.table_name},
        )

        table.grant_read_write_data(endpoint_lambda)

        api = apigw.LambdaRestApi(
            self,
            "graphql_go_api",
            handler=endpoint_lambda,
            proxy=False,
            default_cors_preflight_options=apigw.CorsOptions(
                allow_origins=apigw.Cors.ALL_ORIGINS, allow_methods=apigw.Cors.ALL_METHODS
            ),
        )

        items = api.root.add_resource("graphql")
        items.add_method("POST")  # POST /sandbox
