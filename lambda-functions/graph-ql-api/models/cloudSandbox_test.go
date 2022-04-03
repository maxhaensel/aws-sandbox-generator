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
				Metadata: models.SandboxMetadata{
					Id: "Test",
				},
				AccountName: "Test",
			},
		},
		{
			models.AwsSandbox{
				Metadata: models.SandboxMetadata{
					Id:            "Test",
					AssignedSince: "Test",
					AssignedTo:    "Test",
					AssignedUntil: "Test",
				},
				AccountName: "Test",
			},
		},
		{
			models.AwsSandbox{
				Metadata: models.SandboxMetadata{
					Id:            "",
					AssignedSince: "",
					AssignedTo:    "",
					AssignedUntil: "",
				},
				AccountName: "",
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("error in item %d expect", i+1), func(t *testing.T) {

			if tt.Metadata().Id() != tt.U.Metadata.Id {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.AccountName() != tt.U.AccountName {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.Metadata().AssignedSince() != tt.U.Metadata.AssignedSince {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.Metadata().AssignedTo() != tt.U.Metadata.AssignedTo {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.Metadata().AssignedUntil() != tt.U.Metadata.AssignedUntil {
				t.Errorf("error in item %d expect", i+1)
			}
		})
	}

}

func TestLeaseAzureResolver(t *testing.T) {
	tests := []models.LeaseAzureResolver{
		{
			models.AzureSandbox{
				Metadata: models.SandboxMetadata{
					Id: "Test",
				},
				SandboxName: "Test",
			},
		},
		{
			models.AzureSandbox{
				Metadata: models.SandboxMetadata{
					Id:            "Test",
					AssignedSince: "Test",
					AssignedTo:    "Test",
					AssignedUntil: "Test",
				},
				SandboxName: "Test",
			},
		},
		{
			models.AzureSandbox{
				Metadata: models.SandboxMetadata{
					Id:            "",
					AssignedSince: "",
					AssignedTo:    "",
					AssignedUntil: "",
				},
				SandboxName: "",
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("error in item %d expect", i+1), func(t *testing.T) {

			if tt.Metadata().Id() != tt.U.Metadata.Id {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.SandboxName() != tt.U.SandboxName {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.Metadata().AssignedSince() != tt.U.Metadata.AssignedSince {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.Metadata().AssignedTo() != tt.U.Metadata.AssignedTo {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.Metadata().AssignedUntil() != tt.U.Metadata.AssignedUntil {
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
