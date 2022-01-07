package resolver

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (*Resolver) HelloWorld(ctx context.Context) string {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-central-1"))

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)

	input := &dynamodb.ScanInput{
		TableName: aws.String("AWSSandbox-TableCD117FA1-GIBW29BSQT2O"),
	}
	resp, err := svc.Scan(ctx, input)
	// Build the request with its input parameters
	// resp, err := svc.ListTables(ctx, &dynamodb.ListTablesInput{
	// 	Limit: aws.Int32(5),
	// })
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	fmt.Println("Tables:")
	for _, tableName := range resp.Items {
		x := tableName["account_id"]
		fmt.Println(x)
	}
	return "Hello, world!"
}
