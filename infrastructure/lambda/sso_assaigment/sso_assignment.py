from clear_bucket_policy import clear_bucket_policy
from api.aws_cdk_clients import sso_api, s3_api
from utils.find_user import find_user
from utils.cloud_nuke import cloud_nuke
from utils.handle_sso import handle_sso


def create_sso_assigment(assigned_to: str, account_id: str):
    role = f"arn:aws:iam::172920935848:role/prod-ssocrossrolelessosandboxrole98f4c9f4c18ae60959d1"

    try:
        identitystore_client, sso_admin_client = sso_api(role)
    except:
        return

    user = find_user(assigned_to=assigned_to, identitystore_client=identitystore_client)
    if user == None:
        return

    handle_sso(account_id, user, sso_admin_client, "create")
    # update_database_status()


def remove_sso_assigment_and_nuke(assigned_to: str, account_id: str):
    role = f"arn:aws:iam::172920935848:role/prod-ssocrossrolelessosandboxrole98f4c9f4c18ae60959d1"
    role2 = f"arn:aws:iam::815837829183:role/prod-sandbox-4dboxrolesandbox415420752fec06db2dea9"
    s3_client, credentials_account_to_nuke = s3_api(role2)
    identitystore_client, sso_admin_client = sso_api(role)

    user = find_user(assigned_to=assigned_to, identitystore_client=identitystore_client)
    if user == None:
        return

    # handle_sso(account_id, user, sso_admin_client, "delete")

    # clear_bucket_policy(s3_client)
    cloud_nuke(credentials_account_to_nuke)

    # update_database_status()
