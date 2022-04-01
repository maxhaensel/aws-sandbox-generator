package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/api"
	"lambda/aws-sandbox/graph-ql-api/connection"
	"lambda/aws-sandbox/graph-ql-api/models"
	"lambda/aws-sandbox/graph-ql-api/utils"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/graph-gophers/graphql-go"
)

var valid bool

func (*Resolver) LeaseSandBox(ctx context.Context, args struct {
	Email     string
	LeaseTime string
	Cloud     string
}) (*models.LeaseSandBoxResult, error) {

	valid = utils.ProofPexonMail(args.Email)
	if !valid {
		// 🤦‍♀️
		return nil, fmt.Errorf("no valid Pexon-Mail")

	}

	valid = utils.Lease_time_Input(args.LeaseTime)
	if !valid {
		// 🤦‍♀️
		return nil, fmt.Errorf("Lease-Time is not correct")
	}

	s := strings.Split(args.LeaseTime, "-")
	year, _ := strconv.Atoi(s[0])
	month, _ := strconv.Atoi(s[1])
	day, _ := strconv.Atoi(s[2])

	// check if the Cloud is AZURE
	if args.Cloud == models.PublicCloud.GetAZURE() {
		// do your logic here 🤡
		since, until := utils.TimeRange(year, month, day)
		state_name := strings.Replace(strings.Split(args.Email, "@")[0], ".", "-", 1)

		data := url.Values{
			"rg_name":       {"bootcamp-" + state_name},
			"trainee_email": {args.Email},
			"needed_until":  {*until},
			"created_by":    {args.Email},
		}

		res := models.GitlabPipelineResponse{}
		url := os.Getenv("gitlab_azure_pipeline_webhook")
		url += "&variables[TF_STATE_NAME]=" + state_name

		resp, err := http.PostForm(url, data)
		if err != nil {
			log.Fatal(err)
		}
		json.NewDecoder(resp.Body).Decode(&res)

		return &models.LeaseSandBoxResult{
			Result: &models.LeaseAzureResolver{
				U: models.AzureSandbox{
					Id:            graphql.ID(uuid.New().String()),
					AssignedUntil: *until,
					AssignedSince: *since,
					AssignedTo:    args.Email,
					PipelineId:    strconv.Itoa(res.Id),
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
					Id:            "uuid!",
					AssignedUntil: updatedSandbox.Assigned_until,
					AssignedSince: updatedSandbox.Assigned_since,
					AssignedTo:    updatedSandbox.Assigned_to,
					AccountName:   updatedSandbox.Account_name,
				},
			},
		}, nil

	}

	// TODO 🔥 add custom error to be more clear whats going on
	return nil, fmt.Errorf("internal servererror")
}
