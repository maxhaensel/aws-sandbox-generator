from variables import IdentityStoreId
import logging


def find_user(assigned_to: str, identitystore_client):
    try:
        find_user = identitystore_client.list_users(
            IdentityStoreId=IdentityStoreId,
            Filters=[{"AttributePath": "UserName", "AttributeValue": assigned_to}],
        )
    except Exception:
        logging.ERROR(Exception)

    try:
        if len(find_user["Users"]) == 1:
            return find_user["Users"][0]
    except KeyError:
        logging.ERROR(KeyError)

    return None
