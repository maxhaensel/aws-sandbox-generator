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

	item_available_false_string_1 := models.SandboxItem{
		Account_id: "123",
		Available:  "not-a-bool",
	}

	item_available_false_string_2 := models.SandboxItem{
		Account_id: "123",
		Available:  "123",
	}

	var tests = []struct {
		testname string
		items    []models.SandboxItem
		result   bool
	}{
		{"test1", []models.SandboxItem{item_available, item_available}, true},
		{"test2", []models.SandboxItem{item_available, item_not_available}, true},
		{"test3", []models.SandboxItem{item_not_available, item_not_available}, false},
		{"test4", []models.SandboxItem{item_available_false_string_1}, false},
		{"test5", []models.SandboxItem{item_available_false_string_2}, false},
		{"test5", []models.SandboxItem{}, false},
		{"test5", []models.SandboxItem{}, false},
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
