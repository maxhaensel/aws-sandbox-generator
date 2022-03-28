package resolver

import (
	"context"
	"lambda/aws-sandbox/graph-ql-api/models"
)

// var valid bool

func (*Resolver) LeaseASandBox(ctx context.Context, args struct {
	Email      string
	Lease_time string
	Cloud      string
}) ([]*models.SearchResult, error) {

	// fmt.Println(args.Cloud)

	var l []*models.SearchResult

	l = append(l, &models.SearchResult{Result: &models.AZUREResolver{U: models.AZURE{Pipeline_id: "this-is-azure"}}})

	l = append(l, &models.SearchResult{Result: &models.AWSResolver{U: models.AWS{
		Account_id: "this-is-aws",
	}}})

	return l, nil

	// valid = utils.ProofPexonMail(args.Email)
	// if !valid {
	// 	return nil, fmt.Errorf("no valid Pexon-Mail")

	// }

	// valid = utils.Lease_time_Input(args.Lease_time)
	// if !valid {
	// 	return nil, fmt.Errorf("Lease-Time is not correct")
	// }

	// s := strings.Split(args.Lease_time, "-")
	// year, _ := strconv.Atoi(s[0])
	// month, _ := strconv.Atoi(s[1])
	// day, _ := strconv.Atoi(s[2])

	// svc := connection.GetDynamoDbClient(ctx)

	// items := api.ScanSandboxTable(ctx, svc)

	// sandbox, err := utils.FindAvailableSandbox(items)

	// if err != nil {
	// 	return nil, fmt.Errorf("error while finding a sandbox")
	// }

	// // if sandbox == nil {
	// // 	return &models.SandBoxResolver{U: models.SandBoxResponse{
	// // 		Message: "no Sandbox Available",
	// // 	}}, nil
	// // }

	// since, until := utils.TimeRange(year, month, day)

	// sandbox.Assigned_to = args.Email
	// sandbox.Assigned_since = *since
	// sandbox.Assigned_until = *until
	// sandbox.Available = "false"

	// updatedSandbox, err := api.UpdateSandBoxItem(ctx, svc, *sandbox)

	// fmt.Print(updatedSandbox)

	// if err != nil {
	// 	return nil, err
	// }

	// return &models.SandBoxResolver{U: models.SandBoxResponse{
	// 	Message: "Sandbox is provided",
	// 	Sandbox: models.SandBoxItemResolver{
	// 		U: *updatedSandbox,
	// 	},
	// }}, nil
}
