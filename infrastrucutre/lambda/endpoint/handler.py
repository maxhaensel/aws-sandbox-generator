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


def internal_server_error(msg):
    logger.error(f"{msg}")
    return {
        "statusCode": 500,
        "headers": {"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Credentials": True},
        "body": json.dumps({"message": "Internal Server Error"}),
    }


def handler(event, context):
    try:
        duration_of_lease_in_days = int(os.getenv("duration_of_lease_in_days", 0))
    except:
        return internal_server_error("error")
    try:
        pexonian = event["queryStringParameters"]["name"]
        # format: 2022-01-07
        lease_time = event["queryStringParameters"]["lease_time"]
        year, month, day = lease_time.split("-")
    except:
        return internal_server_error(f"no queryStringParameters provided")

    check_if_mail_is_valid = re.search("\w+\.\w+\@pexon-consulting\.de", pexonian)

    if not check_if_mail_is_valid:
        return internal_server_error(f"someone trys to use sandbox with the following mail: {pexonian} is not allowed")

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
                ":assigned_until": datetime(int(year), int(month), int(day)).isoformat(),
            },
        )
        return {
            "statusCode": 200,
            "headers": {"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Credentials": True},
            "body": json.dumps({"message": "Sandbox is provided", "sandbox": less_than_zero[0]["account_name"]}),
        }
    else:
        return {
            "statusCode": 200,
            "headers": {"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Credentials": True},
            "body": json.dumps({"message": "no sandbox available, come back later!"}),
        }
