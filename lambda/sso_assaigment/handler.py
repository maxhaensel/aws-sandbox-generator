# dummy_event = {
#     "Records": [
#         {
#             "eventID": "40b88af852e92ac246cc2b530430c4b1",
#             "eventName": "MODIFY",
#             "eventVersion": "1.1",
#             "eventSource": "aws:dynamodb",
#             "awsRegion": "eu-central-1",
#             "dynamodb": {
#                 "ApproximateCreationDateTime": 1637685771.0,
#                 "Keys": {"account_id": {"S": "789"}},
#                 "NewImage": {
#                     "account_id": {"S": "807684576972"},
#                     "assigned_since": {"S": "2021-11-23T16:42:51.221223"},
#                     "account_name": {"S": "sandbox-3"},
#                     "available": {"S": "false"},
#                     "assigned_to": {"S": "maximilian.haensel@pexon-consulting.de"},
#                 },
#                 "OldImage": {
#                     "account_id": {"S": "789"},
#                     "assigned_since": {"S": ""},
#                     "account_name": {"S": "sandbox-3"},
#                     "available": {"S": "true"},
#                     "assigned_to": {"S": "test2"},
#                 },
#                 "SequenceNumber": "1400000000036128654773",
#                 "SizeBytes": 201,
#                 "StreamViewType": "NEW_AND_OLD_IMAGES",
#             },
#             "eventSourceARN": "arn:aws:dynamodb:eu-central-1:172920935848:table/AWSSandbox-TableCD117FA1-GIBW29BSQT2O/stream/2021-11-23T16:24:44.525",
#         }
#     ]
# }
import boto3
import subprocess
import os
import json

identitystore_client = boto3.client("identitystore")
sso_admin_client = boto3.client("sso-admin")
client = boto3.client("sts")


def handler(event=None, context=None):
    data_new_image, data_old_image = (
        event["Records"][0]["dynamodb"]["NewImage"],
        event["Records"][0]["dynamodb"]["OldImage"],
    )

    available_new, available_old = data_new_image["available"]["S"], data_old_image["available"]["S"]

    if available_new == "false" and available_old == "true":

        account_id = data_new_image["account_id"]["S"]
        assigned_to = data_new_image["assigned_to"]["S"]

        find_user = identitystore_client.list_users(
            IdentityStoreId="d-99672b9ab3",
            Filters=[{"AttributePath": "UserName", "AttributeValue": assigned_to}],
        )
        if len(find_user["Users"]) > 0:
            sso_admin_client.create_account_assignment(
                InstanceArn="arn:aws:sso:::instance/ssoins-6987b5e3ca99dac9",
                TargetId=account_id,
                TargetType="AWS_ACCOUNT",
                PermissionSetArn="arn:aws:sso:::permissionSet/ssoins-6987b5e3ca99dac9/ps-17b5cee28043a210",
                PrincipalType="USER",
                PrincipalId=find_user["Users"][0]["UserId"],
            )

    if available_new == "true" and available_old == "false":

        account_id = data_old_image["account_id"]["S"]
        assigned_to = data_old_image["assigned_to"]["S"]

        find_user = identitystore_client.list_users(
            IdentityStoreId="d-99672b9ab3",
            Filters=[{"AttributePath": "UserName", "AttributeValue": assigned_to}],
        )

        sso_admin_client.delete_account_assignment(
            InstanceArn="arn:aws:sso:::instance/ssoins-6987b5e3ca99dac9",
            TargetId=account_id,
            TargetType="AWS_ACCOUNT",
            PermissionSetArn="arn:aws:sso:::permissionSet/ssoins-6987b5e3ca99dac9/ps-17b5cee28043a210",
            PrincipalType="USER",
            PrincipalId=find_user["Users"][0]["UserId"],
        )

        assume_role = client.assume_role(
            RoleArn=f"arn:aws:iam::{account_id}:role/OrganizationAccountAccessRole",
            RoleSessionName="cloud-nuke-lambda",
        )

        lambda_key = os.getenv("AWS_ACCESS_KEY_ID")

        credentials = assume_role["Credentials"]

        env = {
            "AWS_ACCESS_KEY_ID": credentials["AccessKeyId"],
            "AWS_SECRET_ACCESS_KEY": credentials["SecretAccessKey"],
            "AWS_SESSION_TOKEN": credentials["SessionToken"],
        }

        if bool(json.loads(os.getenv("dry_run").lower())):
            command = "--dry-run"
        else:
            command = "--force"

        if lambda_key != env["AWS_ACCESS_KEY_ID"]:
            subprocess.check_call([f"./cloud-nuke aws {command}"], shell=True, env=env)

    return {"statusCode": 200}
