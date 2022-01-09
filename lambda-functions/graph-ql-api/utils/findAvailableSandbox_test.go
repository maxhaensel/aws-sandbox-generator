package utils_test

import (
	"lambda/aws-sandbox/graph-ql-api/models"
	"lambda/aws-sandbox/graph-ql-api/utils"
	"testing"
)

func TestFindAvailableSandbox(t *testing.T) {
	item_available := models.SandboxItem{
		Account_id: "123",
		Available:  "true",
	}
	item_not_available := models.SandboxItem{
		Account_id: "456",
		Available:  "false",
	}

	var tests = []struct {
		testname string
		items    []models.SandboxItem
		result   bool
	}{
		{"test1", []models.SandboxItem{item_available, item_available}, true},
		{"test2", []models.SandboxItem{item_available, item_not_available}, true},
		{"test3", []models.SandboxItem{item_not_available, item_not_available}, false},
	}

	for i, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			_, Available, _ := utils.FindAvailableSandbox(tt.items)
			if Available != tt.result {
				t.Errorf("error in item %d expect %t got %t", i+1, Available, tt.result)
			}
		})
	}
}
