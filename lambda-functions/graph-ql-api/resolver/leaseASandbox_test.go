package resolver_test

import (
	"context"
	"fmt"
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
	"github.com/graph-gophers/graphql-go/errors"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

var rootSchema *graphql.Schema
var no_scan_result api.MockedDynamoDB
var no_available_sandbox api.MockedDynamoDB

func init() {
	rootSchema = graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})
	os.Setenv("env", "test")
	os.Setenv("dynamodb_table", "test")

	no_scan_result = api.MockedDynamoDB{
		Scan_response: &dynamodb.ScanOutput{
			Count:            0,
			Items:            []map[string]types.AttributeValue{},
			LastEvaluatedKey: map[string]types.AttributeValue{},
			ScannedCount:     0,
			ResultMetadata:   middleware.Metadata{},
		},
		Scan_err: nil,
	}

	no_available_sandbox = api.MockedDynamoDB{
		Scan_response: &dynamodb.ScanOutput{
			Count: 2,
			Items: []map[string]types.AttributeValue{
				{
					"available": &types.AttributeValueMemberS{Value: "false"},
				},
				{
					"available": &types.AttributeValueMemberS{Value: "false"},
				},
			},
			LastEvaluatedKey: map[string]types.AttributeValue{},
			ScannedCount:     2,
			ResultMetadata:   middleware.Metadata{},
		},
		Scan_err: nil,
	}
}

var malformed_available = api.MockedDynamoDB{
	Scan_response: &dynamodb.ScanOutput{
		Count: 2,
		Items: []map[string]types.AttributeValue{
			{
				"available": &types.AttributeValueMemberS{Value: "xxz"},
			},
		},
		LastEvaluatedKey: map[string]types.AttributeValue{},
		ScannedCount:     2,
		ResultMetadata:   middleware.Metadata{},
	},
	Scan_err: nil,
}

var Query = `
			mutation CreateReviewForEpisode($Email: String!, $Lease_time: String!) {
				leaseASandBox(Email: $Email, Lease_time: $Lease_time) {
					message
				}
			}
	`

var noValideMail = []*errors.QueryError{{
	ResolverError: fmt.Errorf(`no valid Pexon-Mail`),
	Message:       `no valid Pexon-Mail`,
	Path:          []interface{}{"leaseASandBox"}}}

var wrongLeaseTime = []*errors.QueryError{{
	ResolverError: fmt.Errorf(`Lease-Time is not correct`),
	Message:       `Lease-Time is not correct`,
	Path:          []interface{}{"leaseASandBox"}}}

var noSandboxAvailable = `
	{
		"leaseASandBox": {
			"message": "no Sandbox Available"
		}
	}
	`

func TestLeaseASandbox_malformed_input(t *testing.T) {

	tests := []struct {
		svc            api.MockedDynamoDB
		variables      map[string]interface{}
		ExpectedErrors []*errors.QueryError
	}{
		{
			svc: no_scan_result,
			variables: map[string]interface{}{
				"Email":      "party@gmx.de",
				"Lease_time": "2024-05-02",
			},
			ExpectedErrors: noValideMail,
		},
		{
			svc: no_scan_result,
			variables: map[string]interface{}{
				"Email":      "test.test@pexon-consulting.de",
				"Lease_time": "2024",
			},
			ExpectedErrors: wrongLeaseTime,
		},
		{
			svc: malformed_available,
			variables: map[string]interface{}{
				"Email":      "test.test@pexon-consulting.de",
				"Lease_time": "2024-10-20",
			},
			ExpectedErrors: []*errors.QueryError{{
				ResolverError: fmt.Errorf(`error while finding a sandbox`),
				Message:       `error while finding a sandbox`,
				Path:          []interface{}{"leaseASandBox"}}},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			ctx := context.TODO()
			ctx = context.WithValue(ctx, utils.SvcClient, test.svc)
			gqltesting.RunTest(t, &gqltesting.Test{
				Context:        ctx,
				Schema:         rootSchema,
				Variables:      test.variables,
				Query:          Query,
				ExpectedResult: `{"leaseASandBox":null}`,
				ExpectedErrors: test.ExpectedErrors,
			})
		})
	}
}

func TestLeaseASandbox_no_Sandbox_Available(t *testing.T) {

	tests := []struct {
		svc            api.MockedDynamoDB
		expectedResult string
		variables      map[string]interface{}
	}{
		{
			svc:            no_scan_result,
			expectedResult: noSandboxAvailable,
			variables: map[string]interface{}{
				"Email":      "test.test@pexon-consulting.de",
				"Lease_time": "2024-05-02",
			},
		}, {
			svc:            no_available_sandbox,
			expectedResult: noSandboxAvailable,
			variables: map[string]interface{}{
				"Email":      "test.test@pexon-consulting.de",
				"Lease_time": "2024-05-02",
			},
		}}

	for i, test := range tests {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			ctx := context.TODO()
			ctx = context.WithValue(ctx, utils.SvcClient, test.svc)

			gqltesting.RunTests(t, []*gqltesting.Test{
				{
					Context:        ctx,
					Schema:         rootSchema,
					Variables:      test.variables,
					Query:          Query,
					ExpectedResult: test.expectedResult,
				},
			})

		})
	}

}

func TestLeaseASandbox_Scan_Result_not_available(t *testing.T) {

	svc := api.MockedDynamoDB{
		Scan_response: &dynamodb.ScanOutput{
			Count: 2,
			Items: []map[string]types.AttributeValue{
				{
					"account_id":     &types.AttributeValueMemberS{Value: "123"},
					"assigned_to":    &types.AttributeValueMemberS{Value: ""},
					"assigned_since": &types.AttributeValueMemberS{Value: ""},
					"assigned_until": &types.AttributeValueMemberS{Value: ""},
					"available":      &types.AttributeValueMemberS{Value: "true"},
				},
				{
					"account_id":     &types.AttributeValueMemberS{Value: "456"},
					"assigned_to":    &types.AttributeValueMemberS{Value: ""},
					"assigned_since": &types.AttributeValueMemberS{Value: ""},
					"assigned_until": &types.AttributeValueMemberS{Value: ""},
					"available":      &types.AttributeValueMemberS{Value: "true"},
				},
			},
			LastEvaluatedKey: map[string]types.AttributeValue{},
			ScannedCount:     2,
			ResultMetadata:   middleware.Metadata{},
		},
		Scan_err: nil,
		UpdateItem_response: &dynamodb.UpdateItemOutput{
			Attributes: map[string]types.AttributeValue{
				"account_id":     &types.AttributeValueMemberS{Value: "123"},
				"assigned_to":    &types.AttributeValueMemberS{Value: "test.test@pexon-consulting.de"},
				"assigned_since": &types.AttributeValueMemberS{Value: "2022"},
				"assigned_until": &types.AttributeValueMemberS{Value: "2022"},
				"available":      &types.AttributeValueMemberS{Value: "false"},
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
				"Email":      "test.test@pexon-consulting.de",
				"Lease_time": "2024-05-02",
			},
			Query: `
			mutation CreateReviewForEpisode($Email: String!, $Lease_time: String!) {
				leaseASandBox(Email: $Email, Lease_time: $Lease_time) {
						message
						sandbox {
							account_id
							assigned_to
						}
					}
				}
			`,
			ExpectedResult: `{
				"leaseASandBox": {
					"message":"Sandbox is provided",
					"sandbox":{ 
						"account_id" : "123",
						"assigned_to" : "test.test@pexon-consulting.de"
					}
				}
			}
			`,
		},
	})
}
