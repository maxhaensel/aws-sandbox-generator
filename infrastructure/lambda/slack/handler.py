#!/usr/bin/python3.6
import urllib3
import json
import os

http = urllib3.PoolManager()

slack_webhook = os.getenv("SLACK_WEBHOOK")


def handler(event, context):
    msg = {"Content": event["Records"][0]["Sns"]["Message"]}
    payload = {
        "blocks": [
            {"type": "header", "text": {"type": "plain_text", "text": "🚨 error occurred ⚠️", "emoji": True}},
            {"type": "section", "fields": [{"type": "mrkdwn", "text": "*Type:*\nALARM"}]},
            {"type": "section", "fields": [{"type": "mrkdwn", "text": f"*When:*\n{msg['StateChangeTime']}"}]},
            {
                "type": "section",
                "fields": [
                    {"type": "mrkdwn", "text": f"*Function:*\nAWSSandbox-SSOhandlerlambda0DE9C552-AgPpx6Dlp9cT"}
                ],
            },
        ]
    }
    url = slack_webhook
    http.request("POST", url, body=json.dumps(payload).encode("utf-8"))
