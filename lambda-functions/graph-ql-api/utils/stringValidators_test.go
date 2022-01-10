package utils_test

import (
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/utils"
	"testing"
)

func TestProofPexonMail(t *testing.T) {
	tests := []struct {
		value  string
		result bool
	}{
		{"test@mail.de", false},
		{"some_string", false},
		{"test@pexon-consulting.de", false},
		{"test.test@pexon-consulting.de", true},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("check mail: %s ", tt.value), func(t *testing.T) {
			result := utils.ProofPexonMail(tt.value)
			if result != tt.result {
				t.Errorf("error in item %d expect %t got %t", i+1, result, tt.result)
			}
		})
	}

}

func TestLease_time_Input(t *testing.T) {
	tests := []struct {
		value  string
		result bool
	}{
		{"some_string", false},
		{"12313123", false},
		{"2021-13-02", false},
		{"2021-12-32", false},
		{"2021-02-02", true},
		{"2022-12-12", true},
		{"2023-11-11", true},
		{"2024-05-02", true},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("check mail: %s ", tt.value), func(t *testing.T) {
			result := utils.Lease_time_Input(tt.value)
			if result != tt.result {
				t.Errorf("error in item %d expect %t got %t", i+1, result, tt.result)
			}
		})
	}

}
