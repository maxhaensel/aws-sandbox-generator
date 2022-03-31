package models

import "encoding/json"

type GitlabPipelineResponse struct {
	Id            json("id")
	Status
	ProjectId
	WebUrl 
}