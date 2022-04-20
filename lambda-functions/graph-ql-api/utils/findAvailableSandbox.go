package utils

import (
	"lambda/aws-sandbox/graph-ql-api/models"
	"log"
	"strconv"
)

func FindAvailableSandbox(items []*models.LeaseSandBoxResult) (*models.LeaseAwsResolver, error) {

	for _, item := range items {
		if aws_sandbox, ok := item.ToAwsSandbox(); ok {
			b, err := strconv.ParseBool(aws_sandbox.Available())
			if err != nil {
				log.Println(err.Error())
				return nil, err
			}
			if b {
				return aws_sandbox, nil
			}
		}

	}
	return nil, nil
}
