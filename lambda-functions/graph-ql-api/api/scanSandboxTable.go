package api

import (
	"context"
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/models"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func ScanSandboxTable(ctx context.Context, svc DynamoAPI) []models.SandboxItem {

	items := []models.SandboxItem{}

	table := os.Getenv("dynamodb_table")

	if len(table) == 0 {
		err := fmt.Errorf("env-variable dynamodb_table is empty")
		log.Print(fmt.Errorf("ERROR: failed to find table-name %v", err))
		return items
	}

	scanInput := dynamodb.ScanInput{
		TableName: aws.String(table),
	}

	scan, err := svc.Scan(ctx, &scanInput)

	if err != nil {
		log.Print(fmt.Errorf("ERROR: failed to Scan DynamoDB %v", err))
		return items
	}

	err = attributevalue.UnmarshalListOfMaps(scan.Items, &items)

	if err != nil {
		log.Print(fmt.Errorf("ERROR: failed to unmarshal Dynamodb Scan Items, %v", err))
		return items
	}
	return items
}
