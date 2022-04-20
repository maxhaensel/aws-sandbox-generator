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
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func UpdateSandBoxItem(ctx context.Context, svc DynamoAPI, sandbox *models.LeaseSandBoxResult) (*models.LeaseSandBoxResult, error) {

	table := os.Getenv("dynamodb_table")

	if len(table) == 0 {
		err := fmt.Errorf("env-variable dynamodb_table is empty")
		log.Print(fmt.Errorf("ERROR: failed to find table-name %v", err))
		return nil, err
	}
	updateExpressionQuery := `
	SET 
	assigned_to = :assigned_to, 
	assigned_since = :assigned_since, 
	assigned_until = :assigned_until, 
	cloud = :cloud,`
	expr, err := attributevalue.MarshalMap(sandbox)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Record, %w", err)
	}
	if aws_sandbox, ok := sandbox.ToAwsSandbox(); ok {
		accountName := aws_sandbox.AccountName()
		if accountName == "" {
			return nil, fmt.Errorf("no Account_id provided")
		}
		updateExpressionQuery += `available = :available`
		expr, err = attributevalue.MarshalMap(aws_sandbox)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal Record, %w", err)
		}
	}

	if az_sandbox, ok := sandbox.ToAzureSandbox(); ok {
		updateExpressionQuery += `
		pipeline_id = :pipeline_id,
		status = :status,
		project_id = :project_id,
		pipeline_url = :pipeline_url,
		name = :name
		`
		expr, err = attributevalue.MarshalMap(az_sandbox)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal Record, %w", err)
		}
	}
	id := sandbox.CloudSandbox.Id()

	key := map[string]types.AttributeValue{
		"ID": &types.AttributeValueMemberS{Value: string(id)},
	}

	updateExpression := aws.String(updateExpressionQuery)

	input := &dynamodb.UpdateItemInput{
		Key:                       key,
		TableName:                 aws.String(table),
		ExpressionAttributeValues: expr,
		ReturnValues:              "ALL_NEW",
		UpdateExpression:          updateExpression,
	}

	response, err := svc.UpdateItem(ctx, input)

	if err != nil {
		return nil, err
	}

	p := models.LeaseSandBoxResult{}
	attributevalue.UnmarshalMap(response.Attributes, &p)

	return &p, nil
}
