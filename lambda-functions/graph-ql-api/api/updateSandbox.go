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

type DynamoAPIUpdate interface {
	UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
}

func UpdateSandBoxItem(ctx context.Context, svc DynamoAPIUpdate, sandbox models.SandboxItem) (*models.SandboxItem, error) {

	table := os.Getenv("dynamodb_table")

	if len(table) == 0 {
		err := fmt.Errorf("env-variable dynamodb_table is empty")
		log.Print(fmt.Errorf("ERROR: failed to find table-name %v", err))
		return nil, err
	}

	// sandboxItem := struct {
	// 	Account_id string `dynamodbav:"account_id"`
	// }{Account_id: "397624546912"}

	// key, err := attributevalue.MarshalMap(sandboxItem.Account_id)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to marshal Record, %w", err)
	// }

	key := map[string]types.AttributeValue{
		"account_id": &types.AttributeValueMemberS{Value: sandbox.Account_id},
	}

	update := struct {
		Assigned_to    string `dynamodbav:":assigned_to"`
		Assigned_since string `dynamodbav:":assigned_since"`
		Assigned_until string `dynamodbav:":assigned_until"`
		Available      string `dynamodbav:":available"`
	}{
		Assigned_to:    sandbox.Assigned_to,
		Assigned_since: sandbox.Assigned_since,
		Assigned_until: sandbox.Assigned_until,
		Available:      sandbox.Available,
	}

	expr, err := attributevalue.MarshalMap(update)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Record, %w", err)
	}

	updateExpression := aws.String(`
		SET 
		assigned_to = :assigned_to, 
		assigned_since = :assigned_since, 
		assigned_until = :assigned_until, 
		available = :available`,
	)

	input := &dynamodb.UpdateItemInput{
		UpdateExpression:          updateExpression,
		TableName:                 aws.String(table),
		ExpressionAttributeValues: expr,
		Key:                       key,
		ReturnValues:              "ALL_NEW",
	}

	response, err := svc.UpdateItem(ctx, input)

	if err != nil {
		return nil, err
	}

	p := models.SandboxItem{}
	attributevalue.UnmarshalMap(response.Attributes, &p)

	return &p, nil
}
