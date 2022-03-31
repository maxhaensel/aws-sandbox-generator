package models

import "encoding/json"

type GitlabPipelineResponse struct {
	Id            json("id")
	PipelineId    string
	AssignedUntil string
	AssignedSince string
	AssignedTo    string
}