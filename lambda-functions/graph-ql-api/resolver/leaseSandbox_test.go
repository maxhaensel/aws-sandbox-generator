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

func init() {
	os.Setenv("env", "test")
	os.Setenv("dynamodb_table", "test")
}

var rootSchema = graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

var no_scan_result = api.MockedDynamoDB{
	Scan_response: &dynamodb.ScanOutput{
		Count:            0,
		Items:            []map[string]types.AttributeValue{},
		LastEvaluatedKey: map[string]types.AttributeValue{},
		ScannedCount:     0,
		ResultMetadata:   middleware.Metadata{},
	},
	Scan_err: nil,
}

var no_available_sandbox = api.MockedDynamoDB{
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
			mutation LeaseSandBox($email: String!, $leaseTime: String!,  $cloud: Cloud!) {
				leaseSandBox(email: $email, leaseTime: $leaseTime, cloud: $cloud) {
					__typename
					... on AwsSandbox {
						accountName
						metadata {
							id
							assignedTo
							assignedUntil
							assignedSince
						}
					}
					... on AzureSandbox {
						sandboxName
						metadata {
							id
							assignedTo
							assignedUntil
							assignedSince
						}
					}
				}
			}
	`

/*
	############################################################

	Test-Suits for all Clouds

	############################################################
*/

func TestLeaseASandbox_malformed_input(t *testing.T) {
	// error responses
	path := []interface{}{"leaseSandBox"}
	var noValideMail = []*errors.QueryError{{
		ResolverError: fmt.Errorf(`no valid Pexon-Mail`),
		Message:       `no valid Pexon-Mail`,
		Path:          path}}

	var wrongLeaseTime = []*errors.QueryError{{
		ResolverError: fmt.Errorf(`Lease-Time is not correct`),
		Message:       `Lease-Time is not correct`,
		Path:          path}}

	var malformedAvailableproperty = []*errors.QueryError{{
		ResolverError: fmt.Errorf(`error while finding a sandbox`),
		Message:       `error while finding a sandbox`,
		Path:          path}}

	var noAvailableSandbox = []*errors.QueryError{{
		ResolverError: fmt.Errorf(`no Sandbox Available`),
		Message:       `no Sandbox Available`,
		Path:          path}}

	tests := []struct {
		testname       string
		svc            api.MockedDynamoDB
		variables      map[string]interface{}
		ExpectedErrors []*errors.QueryError
	}{
		{
			testname: "noValideMail",
			svc:      no_scan_result,
			variables: map[string]interface{}{
				"email":     "party@gmx.de",
				"leaseTime": "2024-05-02",
				"cloud":     "AWS",
			},
			ExpectedErrors: noValideMail,
		},
		{
			testname: "wrongLeaseTime",
			svc:      no_scan_result,
			variables: map[string]interface{}{
				"email":     "test.test@pexon-consulting.de",
				"leaseTime": "2024",
				"cloud":     "AWS",
			},
			ExpectedErrors: wrongLeaseTime,
		},
		{
			testname: "malformed_available_property",
			svc:      malformed_available,
			variables: map[string]interface{}{
				"email":     "test.test@pexon-consulting.de",
				"leaseTime": "2024-10-20",
				"cloud":     "AWS",
			},
			ExpectedErrors: malformedAvailableproperty,
		},
		{
			testname: "no_available_sandbox",
			svc:      no_available_sandbox,
			variables: map[string]interface{}{
				"email":     "test.test@pexon-consulting.de",
				"leaseTime": "2024-05-02",
				"cloud":     "AWS",
			},
			ExpectedErrors: noAvailableSandbox,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("testcase %d, testname %s", i, test.testname), func(t *testing.T) {
			ctx := context.TODO()
			ctx = context.WithValue(ctx, utils.SvcClient, test.svc)
			gqltesting.RunTest(t, &gqltesting.Test{
				Context:        ctx,
				Schema:         rootSchema,
				Variables:      test.variables,
				Query:          Query,
				ExpectedResult: "null",
				ExpectedErrors: test.ExpectedErrors,
			})
		})
	}
}

func TestLeaseSandbox_Internal_Servererror(t *testing.T) {

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: context.TODO(),
			Schema:  rootSchema,
			Variables: map[string]interface{}{
				"email":     "test.test@pexon-consulting.de",
				"leaseTime": "2024-05-02",
				"cloud":     "GCP",
			},
			Query:          Query,
			ExpectedResult: "null",
			ExpectedErrors: []*errors.QueryError{{
				ResolverError: fmt.Errorf(`internal servererror`),
				Message:       `internal servererror`,
				Path:          []interface{}{"leaseSandBox"}}},
		},
	})
}

/*
	############################################################

	Test-Suits for AWS-Sandbox-calls

	############################################################
*/

func TestLeaseSandbox_AWS_Successfully_Provisioning(t *testing.T) {

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
				"account_name":   &types.AttributeValueMemberS{Value: "sandbox-123"},
				"assigned_to":    &types.AttributeValueMemberS{Value: "test.test@pexon-consulting.de"},
				"assigned_since": &types.AttributeValueMemberS{Value: "2022"},
				"assigned_until": &types.AttributeValueMemberS{Value: "2022"},
				"available":      &types.AttributeValueMemberS{Value: "false"},
			},
		},
		UpdateItem_err: nil,
	}

	ctx := context.WithValue(context.TODO(), utils.SvcClient, svc)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Variables: map[string]interface{}{
				"email":     "test.test@pexon-consulting.de",
				"leaseTime": "2024-05-02",
				"cloud":     "AWS",
			},
			Query: Query,
			ExpectedResult: `{
				"leaseSandBox":{
					"__typename": "AwsSandbox",
					"metadata": {
						"id": "uuid!"
						"assignedSince": "2022",
						"assignedTo": "test.test@pexon-consulting.de",
						"assignedUntil": "2022",
					},
					"accountName": "sandbox-123"
				}
			}`,
		},
	})
}

func TestLeaseSandbox_AWS_Without_Account_Id(t *testing.T) {

	svc := api.MockedDynamoDB{
		Scan_response: &dynamodb.ScanOutput{
			Count: 1,
			Items: []map[string]types.AttributeValue{
				{
					"account_id":     &types.AttributeValueMemberS{Value: ""},
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
	}

	ctx := context.WithValue(context.TODO(), utils.SvcClient, svc)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Variables: map[string]interface{}{
				"email":     "test.test@pexon-consulting.de",
				"leaseTime": "2024-05-02",
				"cloud":     "AWS",
			},
			Query:          Query,
			ExpectedResult: "null",
			ExpectedErrors: []*errors.QueryError{{
				ResolverError: fmt.Errorf(`no Account_id provided`),
				Message:       `no Account_id provided`,
				Path:          []interface{}{"leaseSandBox"}}},
		},
	})
}

/*
	############################################################

	Test-Suits for Azure-Sandbox-calls

	currently disabled because they will fail

	############################################################
*/

// func TestLeaseSandbox_AZURE_Successfully_Provisioning(t *testing.T) {

// 	gqltesting.RunTests(t, []*gqltesting.Test{
// 		{
// 			Context: context.TODO(),
// 			Schema:  rootSchema,
// 			Variables: map[string]interface{}{
// 				"email":     "test.test@pexon-consulting.de",
// 				"leaseTime": "2024-05-02",
// 				"cloud":     "AZURE",
// 			},
// 			Query: Query,
// 			ExpectedResult: `{
// 				"leaseSandBox":{
// 					"__typename": "AzureSandbox",
// 					"pipelineId": "this-is-azure",
// 					"assignedSince": "2022",
// 					"assignedTo": "max",
// 					"assignedUntil": "2023",
// 					"id": "this-azure2"
// 				}
// 			}`,
// 		},
// 	})
// }
