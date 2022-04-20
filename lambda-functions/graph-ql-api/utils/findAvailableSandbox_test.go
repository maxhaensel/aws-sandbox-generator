package utils_test

import (
	"lambda/aws-sandbox/graph-ql-api/models"
	"lambda/aws-sandbox/graph-ql-api/utils"
	"testing"
)

func TestFindAvailableSandbox(t *testing.T) {
	aws := models.AwsSandbox{
		Id:            "123",
		Cloud:         models.Cloud{AZURE: models.PublicCloud.AZURE},
		AssignedUntil: "",
		AssignedSince: "",
		AssignedTo:    "",
		AccountName:   "123",
		Available:     "true",
	}

	item_available := &models.LeaseSandBoxResult{
		CloudSandbox: &models.LeaseAwsResolver{
			Aws: aws,
		},
	}

	aws.Available = "false"
	aws.AccountName = "456"
	item_not_available := &models.LeaseSandBoxResult{
		CloudSandbox: &models.LeaseAwsResolver{
			Aws: aws,
		},
	}
	aws.Available = "123"
	aws.AccountName = "not-a-bool"
	item_available_false_string_1 := &models.LeaseSandBoxResult{
		CloudSandbox: &models.LeaseAwsResolver{
			Aws: aws,
		},
	}
	aws.Available = "123"
	aws.AccountName = "123"
	item_available_false_string_2 := &models.LeaseSandBoxResult{
		CloudSandbox: &models.LeaseAwsResolver{
			Aws: aws,
		},
	}

	var tests = []struct {
		testname string
		items    []*models.LeaseSandBoxResult
		result   bool
	}{
		{"test1", []*models.LeaseSandBoxResult{item_available, item_available}, true},
		{"test2", []*models.LeaseSandBoxResult{item_available, item_not_available}, true},
		{"test3", []*models.LeaseSandBoxResult{item_not_available, item_not_available}, false},
		{"test4", []*models.LeaseSandBoxResult{item_available_false_string_1}, false},
		{"test5", []*models.LeaseSandBoxResult{item_available_false_string_2}, false},
		{"test5", []*models.LeaseSandBoxResult{}, false},
		{"test5", []*models.LeaseSandBoxResult{}, false},
	}

	for i, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			sandbox, _ := utils.FindAvailableSandbox(tt.items)

			var k bool = false
			if sandbox != nil {
				k = true
			}

			if k != tt.result {
				t.Errorf("error in item %d expect %v got %v", i+1, *sandbox, tt.result)
			}
		})
	}
}
