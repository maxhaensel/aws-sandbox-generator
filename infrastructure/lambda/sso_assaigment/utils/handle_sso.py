from variables import InstanceArn, PermissionSetArn


def handle_sso(account_id: str, user, sso_admin_client, method: str):

    func = {"create": sso_admin_client.create_account_assignment, "delete": sso_admin_client.delete_account_assignment}

    test = func[method](
        InstanceArn=InstanceArn,
        TargetId=account_id,
        TargetType="AWS_ACCOUNT",
        PermissionSetArn=PermissionSetArn,
        PrincipalType="USER",
        PrincipalId=user["UserId"],
    )
    print(test)
