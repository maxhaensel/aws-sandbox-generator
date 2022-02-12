package resolver

import (
	"context"
	"lambda/aws-sandbox/graph-ql-api/api"
	"lambda/aws-sandbox/graph-ql-api/connection"
	"lambda/aws-sandbox/graph-ql-api/models"
)

func (*Resolver) ListSandboxes(ctx context.Context) (*models.ListSandboxeResponse, error) {

	svc := connection.GetDynamoDbClient(ctx)

	items := api.ScanSandboxTable(ctx, svc)

	sandboxes := []*models.SandBoxItemResolver{}

	for _, item := range items {
		toadd := models.SandBoxItemResolver{
			U: models.SandboxItem{
				Account_id:     item.Account_id,
				Account_name:   item.Account_name,
				Assigned_since: item.Assigned_since,
				Assigned_to:    item.Assigned_to,
				Assigned_until: item.Assigned_until,
				Available:      item.Available,
			},
		}
		sandboxes = append(sandboxes, &toadd)
	}

	return &models.ListSandboxeResponse{U: models.ListSandboxeResolver{
		Sandboxes: sandboxes}}, nil
}
