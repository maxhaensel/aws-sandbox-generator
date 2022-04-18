package models_test

import (
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/models"
	"testing"
)

func TestLeaseAwsResolver(t *testing.T) {
	tests := []models.LeaseAwsResolver{
		{
			models.AwsSandbox{
				Id:          "Test",
				AccountName: "Test",
			},
		},
		{
			models.AwsSandbox{
				Id:            "Test",
				AccountName:   "Test",
				AssignedSince: "Test",
				AssignedTo:    "Test",
				AssignedUntil: "Test",
			},
		},
		{
			models.AwsSandbox{
				Id:            "",
				AccountName:   "",
				AssignedSince: "",
				AssignedTo:    "",
				AssignedUntil: "",
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("error in item %d expect", i+1), func(t *testing.T) {

			if tt.Id() != tt.Aws.Id {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.AccountName() != tt.Aws.AccountName {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.AssignedSince() != tt.Aws.AssignedSince {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.AssignedTo() != tt.Aws.AssignedTo {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.AssignedUntil() != tt.Aws.AssignedUntil {
				t.Errorf("error in item %d expect", i+1)
			}
		})
	}

}

func TestLeaseAzureResolver(t *testing.T) {
	tests := []models.LeaseAzureResolver{
		{
			models.AzureSandbox{
				Id:         "Test",
				PipelineId: "Test",
			},
		},
		{
			models.AzureSandbox{
				Id:            "Test",
				PipelineId:    "Test",
				AssignedSince: "Test",
				AssignedTo:    "Test",
				AssignedUntil: "Test",
			},
		},
		{
			models.AzureSandbox{
				Id:            "",
				PipelineId:    "",
				AssignedSince: "",
				AssignedTo:    "",
				AssignedUntil: "",
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("error in item %d expect", i+1), func(t *testing.T) {

			if tt.Id() != tt.Az.Id {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.PipelineId() != tt.Az.PipelineId {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.AssignedSince() != tt.Az.AssignedSince {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.AssignedTo() != tt.Az.AssignedTo {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.AssignedUntil() != tt.Az.AssignedUntil {
				t.Errorf("error in item %d expect", i+1)
			}
		})
	}

}

func TestCloudEnum(t *testing.T) {
	if models.PublicCloud.GetAWS() != "AWS" {
		t.Error("error in item dont match AWS")
	}
	if models.PublicCloud.GetAZURE() != "AZURE" {
		t.Error("error in item dont match AZURE")
	}
	if models.PublicCloud.AWS != "AWS" {
		t.Error("error in item dont match AWS")
	}
	if models.PublicCloud.AZURE != "AZURE" {
		t.Error("error in item dont match AZURE")
	}
}
