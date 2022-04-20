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
	aws := models.AwsSandbox{
		Id:            "397624546912",
		Cloud:         models.Cloud{AZURE: models.PublicCloud.AZURE},
		AssignedUntil: "asd2",
		AssignedSince: "asd1",
		AssignedTo:    "4",
		AccountName:   "Sandbox-3",
		Available:     "true",
	}

	sandbox := &models.LeaseSandBoxResult{
		CloudSandbox: &models.LeaseAwsResolver{
			Aws: aws,
		},
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

	aws := models.AwsSandbox{
		Id:            "397624546912",
		Cloud:         models.Cloud{AZURE: models.PublicCloud.AZURE},
		AssignedUntil: "asd2",
		AssignedSince: "asd1",
		AssignedTo:    "4",
		AccountName:   "Sandbox-3",
		Available:     "true",
	}

	sandbox := &models.LeaseSandBoxResult{
		CloudSandbox: &models.LeaseAwsResolver{
			Aws: aws,
		},
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

	aws := models.AwsSandbox{
		Id:            "397624546912",
		Cloud:         models.Cloud{AZURE: models.PublicCloud.AZURE},
		AssignedUntil: "asd2",
		AssignedSince: "asd1",
		AssignedTo:    "4",
		AccountName:   "",
		Available:     "true",
	}

	sandbox := &models.LeaseSandBoxResult{
		CloudSandbox: &models.LeaseAwsResolver{
			Aws: aws,
		},
	}
	_, err := api.UpdateSandBoxItem(ctx, svc, sandbox)

	if err.Error() != fmt.Errorf("no Account_id provided").Error() {
		t.Errorf("%v", err)
	}
}
