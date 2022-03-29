package resolver

import (
	"context"
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/api"
	"lambda/aws-sandbox/graph-ql-api/connection"
	"lambda/aws-sandbox/graph-ql-api/models"
	"lambda/aws-sandbox/graph-ql-api/utils"
	"strconv"
	"strings"
)

var valid bool

func (*Resolver) LeaseSandBox(ctx context.Context, args struct {
	Email      string
	Lease_time string
	Cloud      string
}) (*models.LeaseSandBoxResult, error) {

	valid = utils.ProofPexonMail(args.Email)
	if !valid {
		// 🤦‍♀️
		return nil, fmt.Errorf("no valid Pexon-Mail")

	}

	valid = utils.Lease_time_Input(args.Lease_time)
	if !valid {
		// 🤦‍♀️
		return nil, fmt.Errorf("Lease-Time is not correct")
	}

	s := strings.Split(args.Lease_time, "-")
	year, _ := strconv.Atoi(s[0])
	month, _ := strconv.Atoi(s[1])
	day, _ := strconv.Atoi(s[2])

	// check if the Cloud is AZURE
	if args.Cloud == models.PublicCloud.GetAZURE() {
		// do your logic here
		return &models.LeaseSandBoxResult{
			Result: &models.LeaseAzureResolver{
				U: models.AzureSandbox{
					Id:             "this-azure",
					Assigned_until: "2023",
					Assigned_since: "2022",
					Assigned_to:    "max",
					Pipeline_id:    "this-is-azure",
				},
			},
		}, nil
	}

	// check if the Cloud is AWS
	if args.Cloud == models.PublicCloud.GetAWS() {

		svc := connection.GetDynamoDbClient(ctx)

		items := api.ScanSandboxTable(ctx, svc)

		sandbox, err := utils.FindAvailableSandbox(items)

		if err != nil {
			return nil, fmt.Errorf("error while finding a sandbox")
		}

		if sandbox == nil {
			// TODO 🔥 add custom error to be more clear whats going on
			// if there is no Sandbox for AWS-Available it should not look like a server-error
			return nil, fmt.Errorf("no Sandbox Available")
		}

		since, until := utils.TimeRange(year, month, day)

		sandbox.Assigned_to = args.Email
		sandbox.Assigned_since = *since
		sandbox.Assigned_until = *until
		sandbox.Available = "false"

		updatedSandbox, err := api.UpdateSandBoxItem(ctx, svc, *sandbox)

		if err != nil {
			return nil, err
		}

		return &models.LeaseSandBoxResult{
			Result: &models.LeaseAwsResolver{
				U: models.AwsSandbox{
					Id:             "uuid!",
					Assigned_until: updatedSandbox.Assigned_until,
					Assigned_since: updatedSandbox.Assigned_since,
					Assigned_to:    updatedSandbox.Assigned_to,
					Account_name:   updatedSandbox.Account_name,
				},
			},
		}, nil

	}

	// TODO 🔥 add custom error to be more clear whats going on
	return nil, fmt.Errorf("internal servererror")
}
