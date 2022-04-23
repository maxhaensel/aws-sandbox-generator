from constructs import Construct
from aws_cdk import (
    aws_s3 as s3,
    aws_ssm as ssm,
    aws_iam as iam,
    aws_cloudfront as cloudfront,
    aws_cloudfront_origins as origins,
)

from variables import Enviroments


class AWSSandBoxHosting(Construct):
    def __init__(
        self,
        scope: Construct,
        construct_id: str,
        *,
        ssm_sandbox_domain_uri: ssm.StringParameter,
        enviroment: Enviroments,
    ) -> None:
        super().__init__(scope, construct_id)

        cicd_user = iam.User(self, "cicd-user")

        source_bucket = s3.Bucket(self, "hosting-bucket", versioned=False)

        source_bucket.grant_read_write(cicd_user)

        ssm_sandbox_hosting_bucket_name = ssm.StringParameter(
            self,
            "ParameterSandboxHostingBucketName",
            description="Name of the Hosting Bucket",
            parameter_name=f"/{enviroment.value}/sandbox/hostingbucketname",
            string_value=source_bucket.bucket_name,
        )

        ssm_sandbox_hosting_bucket_name.grant_read(cicd_user)
        ssm_sandbox_domain_uri.grant_read(cicd_user)

        cloudfront.Distribution(
            self,
            "cloudfront_distribution",
            default_behavior=cloudfront.BehaviorOptions(origin=origins.S3Origin(source_bucket)),
            default_root_object="index.html",
        )

        # construct export values
        self.bucket_name = source_bucket.bucket_name
