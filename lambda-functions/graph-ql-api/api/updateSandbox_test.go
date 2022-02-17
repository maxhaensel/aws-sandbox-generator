package api_test

import (
	"context"
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/api"
	"lambda/aws-sandbox/graph-ql-api/models"
	"os"
	"testing"
)

func init() {
	os.Setenv("dynamodb_table", "test")
}

func TestUpdateSandboxNoEnv(t *testing.T) {
	os.Unsetenv("dynamodb_table")

	ctx := context.TODO()

	svc := api.MockedDynamoDB{
		UpdateItem_response: nil,
		UpdateItem_err:      fmt.Errorf("some Error"),
	}

	sandbox := models.SandboxItem{
		Account_id:     "397624546912",
		Account_name:   "Sandbox-3",
		Assigned_to:    "4",
		Assigned_since: "asd1",
		Assigned_until: "asd2",
		Available:      "true",
	}

	if _, err := api.UpdateSandBoxItem(ctx, svc, sandbox); err == nil {
		t.Logf("env-var is not set, but there was no error detected")
	}
}

func TestUpdateSandbox(t *testing.T) {

	os.Setenv("dynamodb_table", "test")

	ctx := context.TODO()

	svc := api.MockedDynamoDB{
		UpdateItem_response: nil,
		UpdateItem_err:      fmt.Errorf("some Error"),
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

func TestUpdateSandboxNoAccountId(t *testing.T) {

	os.Setenv("dynamodb_table", "test")

	ctx := context.TODO()

	svc := api.MockedDynamoDB{
		UpdateItem_response: nil,
		UpdateItem_err:      fmt.Errorf("some Error"),
	}

	sandbox := models.SandboxItem{
		Account_id: "",
	}
	_, err := api.UpdateSandBoxItem(ctx, svc, sandbox)

	if err.Error() != fmt.Errorf("no Account_id provided").Error() {
		t.Errorf("%v", err)
	}
}
