package resolver

import (
	"context"
	"lambda/aws-sandbox/graph-ql-api/models"
)

func (*Resolver) ListSandboxes(ctx context.Context) (*models.ListSandboxeResponse, error) {
	return &models.ListSandboxeResponse{U: models.ListSandboxeResolver{
		Sandboxes: []*models.SandBoxItemResolver{
			{
				U: models.SandboxItem{Account_id: "1", Account_name: "4"},
			},
			{
				U: models.SandboxItem{Account_id: "2", Account_name: "5"},
			},
			{
				U: models.SandboxItem{Account_id: "3", Account_name: "6"},
			},
		}}}, nil
}
