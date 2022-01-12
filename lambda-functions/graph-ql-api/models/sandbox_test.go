package models_test

import (
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/models"
	"testing"
)

func TestSandBoxItemResolver(t *testing.T) {
	tests := []models.SandBoxItemResolver{
		{
			models.SandboxItem{
				Account_id:   "Test",
				Account_name: "Test",
			},
		},
		{
			models.SandboxItem{
				Account_id:   "",
				Account_name: "",
			},
		},
		{
			models.SandboxItem{
				Account_id:     "Test",
				Account_name:   "Test",
				Assigned_since: "Test",
				Assigned_to:    "Test",
				Assigned_until: "Test",
				Available:      "Test",
			},
		},
		{
			models.SandboxItem{
				Account_id:     "",
				Account_name:   "",
				Assigned_since: "",
				Assigned_to:    "",
				Assigned_until: "",
				Available:      "",
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("error in item %d expect", i+1), func(t *testing.T) {

			if tt.Account_id() != tt.U.Account_id {
				t.Errorf("error in item %d expect", i+1)
			}

			if tt.Account_name() != tt.U.Account_name {
				t.Errorf("error in item %d expect", i+1)
			}

			if *tt.Assigned_since() != tt.U.Assigned_since {
				t.Errorf("error in item %d expect", i+1)
			}

			if *tt.Assigned_to() != tt.U.Assigned_to {
				t.Errorf("error in item %d expect", i+1)
			}

			if *tt.Assigned_until() != tt.U.Assigned_until {
				t.Errorf("error in item %d expect", i+1)
			}

			if *tt.Available() != tt.U.Available {
				t.Errorf("error in item %d expect", i+1)
			}
		})
	}

}

func TestListSandboxeResponse(t *testing.T) {
	s := models.ListSandboxeResponse{
		models.ListSandboxeResolver{

			Sandboxes: []*models.SandBoxItemResolver{
				{
					models.SandboxItem{
						Account_id:   "Test",
						Account_name: "Test",
					},
				},
				{
					models.SandboxItem{
						Account_id:   "",
						Account_name: "",
					},
				},
			},
		},
	}

	if s.Sandboxes() != &s.U.Sandboxes {
		t.Errorf("error in item")
	}

	if len(*s.Sandboxes()) != 2 {
		t.Errorf("error in item")
	}

	if s.U.Sandboxes[0].Account_id() != s.U.Sandboxes[0].U.Account_id {
		t.Errorf("error in item")
	}

}

func TestSandBoxResolver(t *testing.T) {
	s := models.SandBoxResolver{
		models.SandBoxResponse{
			Message: "test",
			Sandbox: models.SandBoxItemResolver{
				models.SandboxItem{},
			},
		},
	}

	if s.Message() != s.U.Message {
		t.Errorf("error in item")
	}

	if s.Sandbox() != &s.U.Sandbox {
		t.Errorf("error in item")
	}

	if s.Sandbox().Account_id() != s.Sandbox().U.Account_id {
		t.Errorf("error in item")
	}

}

func TestCreateNewSandboxResolver(t *testing.T) {
	s := models.CreateNewSandboxResolver{
		models.SandBoxItemResolver{
			models.SandboxItem{
				Account_id: "123",
			},
		},
	}

	if s.Sandbox.Account_id() != s.Sandbox.U.Account_id {
		t.Errorf("error in item")
	}

	k := models.CreateNewSandboxResponse{}
	k.Sandbox()
}

func TestCreateNewSandboxResponse(t *testing.T) {
	s := models.CreateNewSandboxResponse{
		models.CreateNewSandboxResolver{},
	}

	if s.Sandbox() != &s.U.Sandbox {
		t.Errorf("error in item")
	}

}
