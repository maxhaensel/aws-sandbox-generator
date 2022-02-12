package resolver

import (
	"context"
	"lambda/aws-sandbox/graph-ql-api/models"
)

func (*Resolver) CreateNewSandbox(ctx context.Context, args struct {
	Account_id   string
	Account_name string
}) (*models.CreateNewSandboxResponse, error) {

	return &models.CreateNewSandboxResponse{
		U: models.CreateNewSandboxResolver{
			Sandbox: models.SandBoxItemResolver{
				U: models.SandboxItem{
					Account_id:   "123",
					Account_name: "sandbox.Account_name",
				},
			},
		}}, nil
}
