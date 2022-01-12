package api_test

import (
	"bytes"
	"context"
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/api"
	"lambda/aws-sandbox/graph-ql-api/models"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
)

type mockedScanApi struct {
	response *dynamodb.ScanOutput
	err      error
}

type mockedScanRequest struct {
	svc        api.DynamoAPI
	len        int
	logMessage string
}

func (m mockedScanApi) Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	return m.response, m.err
}

func TestScanSandboxTable(t *testing.T) {

	os.Setenv("dynamodb_table", "test")

	tests := []mockedScanRequest{
		{
			mockedScanApi{
				response: &dynamodb.ScanOutput{
					Count: 2,
					Items: []map[string]types.AttributeValue{
						{
							"account_id":     &types.AttributeValueMemberS{Value: "123"},
							"assigned_to":    &types.AttributeValueMemberS{Value: "123"},
							"assigned_since": &types.AttributeValueMemberS{Value: "123"},
							"assigned_until": &types.AttributeValueMemberS{Value: "123"},
							"available":      &types.AttributeValueMemberS{Value: "true"},
						},
						{
							"account_id":     &types.AttributeValueMemberS{Value: "123"},
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
				err: nil,
			},
			2,
			"",
		},
		{
			mockedScanApi{
				response: &dynamodb.ScanOutput{
					Count: 2,
					Items: []map[string]types.AttributeValue{
						{},
						{},
					},
					LastEvaluatedKey: map[string]types.AttributeValue{},
					ScannedCount:     2,
					ResultMetadata:   middleware.Metadata{},
				},
				err: nil,
			},
			2,
			"",
		},
		{
			mockedScanApi{
				response: nil,
				err:      fmt.Errorf("error"),
			},
			0,
			"ERROR: failed to Scan DynamoDB",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("testcase %d, contain %d items", i, tt.len), func(t *testing.T) {

			ctx := context.TODO()
			items := api.ScanSandboxTable(ctx, tt.svc)

			result := reflect.TypeOf(items) == reflect.TypeOf([]models.SandboxItem{})
			numItems := len(items)

			if !(result == true && numItems == tt.len) {
				t.Errorf("error in item %d expect %d", i+1, tt.len)
			}
		})
	}
}

func TestScanSandboxTableWithoutTableName(t *testing.T) {
	os.Unsetenv("dynamodb_table")

	testcase := mockedScanRequest{
		mockedScanApi{
			response: nil,
			err:      fmt.Errorf("error"),
		},
		0,
		"ERROR: failed to find table-name env-variable dynamodb_table is empty",
	}

	var str bytes.Buffer
	log.SetOutput(&str)

	ctx := context.TODO()
	items := api.ScanSandboxTable(ctx, testcase.svc)

	logMessage := strings.TrimSuffix(str.String(), "\n")

	if !strings.Contains(logMessage, testcase.logMessage) {
		t.Errorf("error in logMessage expect Message got %s", logMessage)
	}

	result := reflect.TypeOf(items) == reflect.TypeOf([]models.SandboxItem{})
	numItems := len(items)

	if !(result == true && numItems == testcase.len) {
		t.Errorf("error in item expect %d", testcase.len)
	}
}
