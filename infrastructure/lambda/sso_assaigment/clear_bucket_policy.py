import boto3
import logging


def clear_bucket_policy(s3):
    # try:
    #     # session = boto3.Session(
    #     #     aws_access_key_id=credentials["AccessKeyId"],
    #     #     aws_secret_access_key=credentials["SecretAccessKey"],
    #     #     aws_session_token=credentials["SessionToken"],
    #     # )
    #     # s3 = session.client("s3")
    #     # logging.debug(session)
    # except:
    #     logging.debug(session)
    #     logging.error("unable to create client for s3")

    try:
        buckets = s3.list_buckets()
        logging.debug(buckets)
    except:
        logging.debug(buckets)
        logging.error("cant list Buckets")

    for bucket in buckets["Buckets"]:
        try:
            response = s3.delete_bucket_policy(Bucket=bucket["Name"])
            logging.debug(response)
        except:
            logging.debug(response)
            logging.error(f"unable to delete bucket_policy for bucket: {bucket['Name']}")
