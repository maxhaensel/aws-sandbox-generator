package resolver_test

import (
	"context"
	"lambda/aws-sandbox/graph-ql-api/api"
	"lambda/aws-sandbox/graph-ql-api/resolver"
	"lambda/aws-sandbox/graph-ql-api/schema"
	"lambda/aws-sandbox/graph-ql-api/utils"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

func TestLeaseASandBox(t *testing.T) {
	os.Setenv("env", "test")
	os.Setenv("dynamodb_table", "test")

	svc := api.MockedDynamoDB{
		Scan_response: &dynamodb.ScanOutput{
			Count: 2,
			Items: []map[string]types.AttributeValue{
				{
					"account_id":     &types.AttributeValueMemberS{Value: "123"},
					"account_name":   &types.AttributeValueMemberS{Value: "name"},
					"assigned_to":    &types.AttributeValueMemberS{Value: "123"},
					"assigned_since": &types.AttributeValueMemberS{Value: "123"},
					"assigned_until": &types.AttributeValueMemberS{Value: "123"},
					"available":      &types.AttributeValueMemberS{Value: "true"},
				},
				{
					"account_id":     &types.AttributeValueMemberS{Value: "123"},
					"account_name":   &types.AttributeValueMemberS{Value: "name"},
					"assigned_to":    &types.AttributeValueMemberS{Value: "123"},
					"assigned_since": &types.AttributeValueMemberS{Value: "123"},
					"assigned_until": &types.AttributeValueMemberS{Value: "123"},
					"available":      &types.AttributeValueMemberS{Value: "true"},
				},
			},
			LastEvaluatedKey: map[string]types.AttributeValue{},
			ScannedCount:     2,
			ResultMetadata:   middleware.Metadata{},
		},
		Scan_err: nil,
	}

	ctx := context.TODO()
	ctx = context.WithValue(ctx, utils.SvcClient, svc)
	rootSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Query: `
				{
					listSandboxes {
						sandboxes {
							account_id
							account_name
						}
					}
				}
			`,
			ExpectedResult: `
				{
					"listSandboxes": {
					  "sandboxes": [
						  {"account_id":"123","account_name":"name"},
						  {"account_id":"123","account_name":"name"}
						]
					}
				}
			`,
		},
	})
}
