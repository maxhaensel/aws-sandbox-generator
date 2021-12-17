import json
import boto3
import os
from datetime import datetime, timedelta
import re
import logging

logger = logging.getLogger()
logger.setLevel(logging.INFO)

dynamodb_table = os.environ.get("dynamodb_table")
table = boto3.resource("dynamodb").Table(dynamodb_table)


def handler(event, context):
    try:
        duration_of_lease_in_days = int(os.getenv("duration_of_lease_in_days", 0))
    except:
        return {"statusCode": 500, "body": json.dumps({"message": "Internal Server Error"})}

    try:
        pexonian = event["queryStringParameters"]["name"]
    except:
        logger.info(f"no queryStringParameters provided")
        return {"statusCode": 500, "body": json.dumps({"message": "Internal Server Error"})}

    check_if_mail_is_valid = re.search("\w+\.\w+\@pexon-consulting\.de", pexonian)

    if not check_if_mail_is_valid:
        logger.info(f"someone trys to use sandbox with the following mail: {pexonian} is not allowed")
        return {"statusCode": 500, "body": json.dumps({"message": "Internal Server Error"})}

    scan = table.scan()
    less_than_zero = list(filter(lambda x: x["available"] == "true", scan["Items"]))

    if len(less_than_zero) != 0:
        table.update_item(
            Key={"account_id": less_than_zero[0]["account_id"]},
            UpdateExpression="SET assigned_to = :assigned_to, available = :available, assigned_since = :assigned_since, assigned_until = :assigned_until",
            ExpressionAttributeValues={
                ":assigned_to": pexonian,
                ":available": "false",
                ":assigned_since": datetime.now().isoformat(),
                ":assigned_until": datetime.now() + timedelta(hours=duration_of_lease_in_days * 24),
            },
        )
        return {
            "statusCode": 200,
            "body": json.dumps({"message": "Sandbox is provided", "sandbox": less_than_zero[0]["account_name"]}),
        }
    else:
        return {"statusCode": 200, "body": json.dumps({"message": "no sandbox available, come back later!"})}
