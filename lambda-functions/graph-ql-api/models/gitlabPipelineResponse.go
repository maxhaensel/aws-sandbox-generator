package models

type GitlabPipelineResponse struct {
	Id        int    `json:"id"`
	Status    string `json:"status"`
	ProjectId int    `json:"project_id"`
	WebUrl    string `json:"web_url"`
}
