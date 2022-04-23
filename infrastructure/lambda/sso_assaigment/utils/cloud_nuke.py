from distutils import util
import json
import logging
import os
import subprocess


def cloud_nuke(credentials) -> bool:
    lambda_key = os.getenv("AWS_ACCESS_KEY_ID")
    env = {
        "AWS_ACCESS_KEY_ID": credentials["AccessKeyId"],
        "AWS_SECRET_ACCESS_KEY": credentials["SecretAccessKey"],
        "AWS_SESSION_TOKEN": credentials["SessionToken"],
    }
    try:
        dry_run = util.strtobool(os.getenv("dry_run", False).lower())
    except Exception as e:
        logging.error(f'Fail to parse env-var "dry_run" not a bool value')
        raise

    command = "--dry-run" if dry_run else "--force"

    try:
        if lambda_key != env["AWS_ACCESS_KEY_ID"]:
            logging.debug("")
            subprocess.check_call(
                [
                    f"./cloud-nuke_darwin_arm64 aws --dry-run --log-level debug --region eu-central-1 --config ./cloud_nuke_config.yaml"
                ],
                shell=True,
                env=env,
            )
            # subprocess.check_call([f"./cloud-nuke aws {command} --log-level error"], shell=True, env=env)
        return True

    except Exception as e:
        logging.error(f"Fail to call cloud-nuke: {e}")
        raise
