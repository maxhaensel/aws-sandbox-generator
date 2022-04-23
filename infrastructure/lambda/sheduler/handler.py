import json
import boto3
import os
from datetime import datetime, timedelta

dynamodb_table = os.environ.get("dynamodb_table")
table = boto3.resource("dynamodb").Table(dynamodb_table)


def handler(event=None, context=None):
    scan = table.scan()

    to_remove_from_sandbox = list(
        filter(lambda x: x["assigned_until"] < datetime.now().isoformat() and x["available"] == "false", scan["Items"])
    )

    for user in to_remove_from_sandbox:
        table.update_item(
            Key={"account_id": user["account_id"]},
            UpdateExpression="SET assigned_to = :assigned_to, available = :available, assigned_since = :assigned_since, assigned_until = :assigned_until",
            ExpressionAttributeValues={
                ":assigned_to": "",
                ":available": "true",
                ":assigned_since": "",
                ":assigned_until": "",
            },
        )

    return {"statusCode": 200}
