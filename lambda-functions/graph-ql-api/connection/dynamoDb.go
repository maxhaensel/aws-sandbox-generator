package connection

import (
	"context"
	"lambda/aws-sandbox/graph-ql-api/api"
	"lambda/aws-sandbox/graph-ql-api/utils"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetDynamoDbClient(ctx context.Context) api.DynamoAPI {
	if os.Getenv("env") == "test" {
		svcUntyped := ctx.Value(utils.SvcClient)
		svc, b := svcUntyped.(api.MockedDynamoDB)
		if b {
			return svc
		} else {
			return api.MockedDynamoDB{}
		}

	}

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-central-1"))

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)

	return svc
}
