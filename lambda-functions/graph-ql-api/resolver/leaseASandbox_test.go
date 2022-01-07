package resolver

import (
	"context"
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/connection"
	"testing"
)

func TestBasic(t *testing.T) {

	resolver := Resolver{}

	x, _ := resolver.LeaseASandBox(context.TODO(), struct {
		Email string
	}{
		Email: "asda",
	})
	fmt.Print(x.u.Sandbox)
}

func TestGetItems(t *testing.T) {
	ctx := context.TODO()
	svc := connection.GetDynamoDbClient(ctx)
	items := getItems(ctx, svc)

	numItems := len(items)

	if numItems == 0 {
		t.Errorf("scan should not retrieve 0 items")
	}
}

func TestGetASandBoxIfPossible(t *testing.T) {
	item_available := SandboxItem{
		Account_id: "123",
		Available:  "true",
	}
	item_not_available := SandboxItem{
		Account_id: "456",
		Available:  "false",
	}

	var tests = []struct {
		testname string
		items    []SandboxItem
		result   bool
	}{
		{"test1", []SandboxItem{item_available, item_available}, true},
		{"test2", []SandboxItem{item_available, item_not_available}, true},
		{"test3", []SandboxItem{item_not_available, item_not_available}, false},
	}

	for i, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			_, Available, _ := getASandBoxIfPossible(tt.items)
			if Available != tt.result {
				t.Errorf("error in item %d expect %t got %t", i+1, Available, tt.result)
			}
		})
	}
}
