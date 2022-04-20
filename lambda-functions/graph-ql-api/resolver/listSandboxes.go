package resolver

import (
	"context"
	"lambda/aws-sandbox/graph-ql-api/api"
	"lambda/aws-sandbox/graph-ql-api/connection"
	"lambda/aws-sandbox/graph-ql-api/models"
)

func (*Resolver) ListSandboxes(ctx context.Context, args struct {
	Email string
}) (*models.ListLeaseSandBoxResult, error) {

	svc := connection.GetDynamoDbClient(ctx)

	items := api.ScanSandboxTable(ctx, svc)

	sandboxes := &models.ListLeaseSandBoxResult{
		CloudSandboxes: items,
	}

	return sandboxes, nil
}
