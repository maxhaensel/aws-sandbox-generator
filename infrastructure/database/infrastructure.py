from ast import Delete
from constructs import Construct
from aws_cdk import (
    aws_s3 as s3,
    aws_ssm as ssm,
    aws_lambda as lambda_,
    aws_iam as iam,
    aws_cloudfront as cloudfront,
    aws_cloudfront_origins as origins,
    aws_apigateway as apigw,
    aws_dynamodb as dynamodb,
    RemovalPolicy,
)

"""
Datamodel
{
    "id": "<uuid>",
    "sandbox_name": "",
    "state#id": "",
    "cloud": "aws | azure | gcp"
    "assigned_until": "2022-04-03T00:00:00Z",
    "assigned_since": "2022-03-02T22:09:52Z",
    "assigned_to": "<email-adresse>",
    "state": "pending | completed | ready",
    "azure": {"pipeline-id": ""},
    "aws": {
        "account_name": "",
        "account_id": "",
    },
}
"""


"""
PrimaryKey: assigned_to
SortKey: id

#### local_secondary_index
SortKey: state


#### global_secondary_index
PrimaryKey: cloud
SortKey: state#id
"""


class AWSTable(Construct):
    def __init__(self, scope: Construct, construct_id: str) -> None:
        super().__init__(scope, construct_id)

        table = dynamodb.Table(
            self,
            "MultiCloudSandbox",
            partition_key=dynamodb.Attribute(name="assigned_to", type=dynamodb.AttributeType.STRING),
            sort_key=dynamodb.Attribute(name="id", type=dynamodb.AttributeType.STRING),
            stream=dynamodb.StreamViewType.NEW_AND_OLD_IMAGES,
            billing_mode=dynamodb.BillingMode.PAY_PER_REQUEST,
            removal_policy=RemovalPolicy.DESTROY,
        )

        table.add_local_secondary_index(
            index_name="lsi_state",
            sort_key=dynamodb.Attribute(name="state", type=dynamodb.AttributeType.STRING),
            projection_type=dynamodb.ProjectionType.INCLUDE,
            non_key_attributes=["sandbox_name", "aws", "azure"],
        )

        table.add_global_secondary_index(
            index_name="gsi_cloud",
            partition_key=dynamodb.Attribute(name="cloud", type=dynamodb.AttributeType.STRING),
            sort_key=dynamodb.Attribute(name="state#id", type=dynamodb.AttributeType.STRING),
            projection_type=dynamodb.ProjectionType.INCLUDE,
            non_key_attributes=["sandbox_name", "aws", "azure", "assigned_to"],
        )

        self.table = table
