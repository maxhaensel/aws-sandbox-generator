package resolver_test

import (
	"context"
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/api"
	"lambda/aws-sandbox/graph-ql-api/resolver"
	"lambda/aws-sandbox/graph-ql-api/schema"
	"lambda/aws-sandbox/graph-ql-api/utils"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/errors"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

func TestDeallocateSandbox(t *testing.T) {
	rootSchema = graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	svc := api.MockedDynamoDB{
		UpdateItem_response: &dynamodb.UpdateItemOutput{
			Attributes: map[string]types.AttributeValue{
				"account_id":     &types.AttributeValueMemberS{Value: "1"},
				"account_name":   &types.AttributeValueMemberS{Value: "account_1"},
				"assigned_to":    &types.AttributeValueMemberS{Value: ""},
				"assigned_since": &types.AttributeValueMemberS{Value: ""},
				"assigned_until": &types.AttributeValueMemberS{Value: ""},
				"available":      &types.AttributeValueMemberS{Value: "true"},
			},
		},
		UpdateItem_err: nil,
	}

	ctx := context.TODO()
	ctx = context.WithValue(ctx, utils.SvcClient, svc)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Variables: map[string]interface{}{
				"Account_id": "",
			},
			Query: `
			mutation RunDeallocateSandbox($Account_id: String!) {
				deallocateSandbox(Account_id: $Account_id) {
						message
						sandbox {
							account_id
							assigned_to
							available
						}
					}
				}
			`,
			ExpectedErrors: []*errors.QueryError{{
				ResolverError: fmt.Errorf(`no Account_id provided`),
				Message:       `no Account_id provided`,
				Path:          []interface{}{"deallocateSandbox"}}},
			ExpectedResult: `{"deallocateSandbox":null}`,
		},
		{
			Context: ctx,
			Schema:  rootSchema,
			Variables: map[string]interface{}{
				"Account_id": "1",
			},
			Query: `
			mutation RunDeallocateSandbox($Account_id: String!) {
				deallocateSandbox(Account_id: $Account_id) {
						message
						sandbox {
							account_id
							assigned_to
							available
						}
					}
				}
			`,
			ExpectedResult: `{
				"deallocateSandbox": {
					"message":"Sandbox successfully deallocate",
					"sandbox":{
						"account_id" : "1",
						"assigned_to" : "",
						"available" : "true"
					}
				}
			}
			`,
		},
	})
}
