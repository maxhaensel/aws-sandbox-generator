package resolver

import (
	"context"
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/connection"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type SandboxItem struct {
	Account_id     string
	Account_name   string
	Assigned_since string
	Assigned_to    string
	Assigned_until string
	Available      string
}

type SandBoxResponse struct {
	Message string
	Sandbox string
}

type SandBoxResolver struct {
	u *SandBoxResponse
}

func (r *SandBoxResolver) Message() *string {
	return &r.u.Message
}

func (r *SandBoxResolver) Sandbox() *string {
	return &r.u.Sandbox
}

func getItems(ctx context.Context, svc *dynamodb.Client) []SandboxItem {

	scanInput := dynamodb.ScanInput{
		//TableName: aws.String("test"),
		TableName: aws.String("AWSSandbox-TableCD117FA1-GIBW29BSQT2O"),
	}

	scan, err := svc.Scan(ctx, &scanInput)

	if err != nil {
		fmt.Print(err)
	}

	items := []SandboxItem{}

	err = attributevalue.UnmarshalListOfMaps(scan.Items, &items)

	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
	}

	return items
}

func getASandBoxIfPossible(items []SandboxItem) (SandboxItem, bool, error) {
	result := []SandboxItem{}

	for _, item := range items {
		b, err := strconv.ParseBool(item.Available)
		if err != nil {
			return SandboxItem{}, false, err
		}
		if b {
			result = append(result, item)
		}
	}

	if len(result) > 0 {
		return result[0], true, nil
	} else {
		return SandboxItem{}, false, nil
	}
}

func (*Resolver) LeaseASandBox(ctx context.Context, args struct {
	Email string
}) (*SandBoxResolver, error) {

	svc := connection.GetDynamoDbClient(ctx)

	items := getItems(ctx, svc)

	sandbox, available, _ := getASandBoxIfPossible(items)

	if !available {
		return &SandBoxResolver{&SandBoxResponse{
			Message: "no Sandbox Available",
			Sandbox: "",
		}}, nil
	}

	return &SandBoxResolver{&SandBoxResponse{
		Message: "Sandbox is provided",
		Sandbox: sandbox.Account_name,
	}}, nil

	// for _, item := range items {
	// 	fmt.Println("Account_id:", item.Account_id)
	// 	fmt.Println("Assigned_since:", item.Assigned_since)
	// 	fmt.Println("Assigned_to:", item.Assigned_to)
	// 	fmt.Println("Assigned_until:", item.Assigned_until)
	// 	fmt.Println("Available:", item.Available)
	// 	fmt.Println()
	// }

	// input := &dynamodb.UpdateItemInput{
	// 	TableName: aws.String("test"),
	// 	Key: map[string]types.AttributeValue{
	// 		"Year": types.AttributeValue{S: "Hello"},
	// 	},
	// }

	// svc.UpdateItem(ctx, input)
	// return args.Email
}
