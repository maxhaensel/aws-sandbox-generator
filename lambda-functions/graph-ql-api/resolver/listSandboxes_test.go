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

	Query := `
			query ListSandboxes($Email: String!)	{
					listSandboxes (Email: $Email) 
					 {
						sandboxes {
							account_id
							account_name
							assigned_to
						}
					}
				}
			`

	svc := api.MockedDynamoDB{
		Scan_response: &dynamodb.ScanOutput{
			Count: 2,
			Items: []map[string]types.AttributeValue{
				{
					"account_id":     &types.AttributeValueMemberS{Value: "1"},
					"account_name":   &types.AttributeValueMemberS{Value: "account_1"},
					"assigned_to":    &types.AttributeValueMemberS{Value: "test.test@pexon-consulting.de"},
					"assigned_since": &types.AttributeValueMemberS{Value: ""},
					"assigned_until": &types.AttributeValueMemberS{Value: ""},
					"available":      &types.AttributeValueMemberS{Value: ""},
				},
				{
					"account_id":     &types.AttributeValueMemberS{Value: "2"},
					"account_name":   &types.AttributeValueMemberS{Value: "account_2"},
					"assigned_to":    &types.AttributeValueMemberS{Value: "some.other@pexon-consulting.de"},
					"assigned_since": &types.AttributeValueMemberS{Value: "123"},
					"assigned_until": &types.AttributeValueMemberS{Value: "123"},
					"available":      &types.AttributeValueMemberS{Value: "true"},
				},
				{
					"account_id":     &types.AttributeValueMemberS{Value: "3"},
					"account_name":   &types.AttributeValueMemberS{Value: "account_3"},
					"assigned_to":    &types.AttributeValueMemberS{Value: "test.test@pexon-consulting.de"},
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
			Variables: map[string]interface{}{
				"Email": "test.test@pexon-consulting.de",
			},
			Query: Query,
			ExpectedResult: `
				{
					"listSandboxes": {
					  "sandboxes": [
						  {"account_id":"1","account_name":"account_1","assigned_to":"test.test@pexon-consulting.de"},
						  {"account_id":"3","account_name":"account_3","assigned_to":"test.test@pexon-consulting.de"}
						]
					}
				}
			`,
		},
		{
			Context: ctx,
			Schema:  rootSchema,
			Variables: map[string]interface{}{
				"Email": "some.other@pexon-consulting.de",
			},
			Query: Query,
			ExpectedResult: `
				{
					"listSandboxes": {
					  "sandboxes": [
						  {"account_id":"2","account_name":"account_2","assigned_to":"some.other@pexon-consulting.de"}
						]
					}
				}
			`,
		},
		{
			Context: ctx,
			Schema:  rootSchema,
			Variables: map[string]interface{}{
				"Email": "not.exist@pexon-consulting.de",
			},
			Query: Query,
			ExpectedResult: `
				{
					"listSandboxes": {
					  "sandboxes": []
					}
				}
			`,
		},
	})
}
