package api

import (
	"context"
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func ScanSandboxTable(ctx context.Context, svc *dynamodb.Client) []models.SandboxItem {

	scanInput := dynamodb.ScanInput{
		TableName: aws.String("test"),
		//TableName: aws.String("AWSSandbox-TableCD117FA1-GIBW29BSQT2O"),
	}

	scan, err := svc.Scan(ctx, &scanInput)

	if err != nil {
		fmt.Print(err)
	}

	items := []models.SandboxItem{}

	err = attributevalue.UnmarshalListOfMaps(scan.Items, &items)

	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
	}

	return items
}
