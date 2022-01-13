package api_test

import (
	"context"
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/api"
	"lambda/aws-sandbox/graph-ql-api/models"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type mockedUpdateApi struct {
	response *dynamodb.UpdateItemOutput
	err      error
}

func (m mockedUpdateApi) UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
	return m.response, m.err
}

func TestUpdateSandbox(t *testing.T) {

	os.Setenv("dynamodb_table", "test")

	ctx := context.TODO()

	svc := mockedUpdateApi{
		response: nil,
		err:      fmt.Errorf("some Error"),
	}

	sandbox := models.SandboxItem{
		Account_id:     "397624546912",
		Account_name:   "Sandbox-3",
		Assigned_to:    "4",
		Assigned_since: "asd1",
		Assigned_until: "asd2",
		Available:      "true",
	}
	api.UpdateSandBoxItem(ctx, svc, sandbox)
}
