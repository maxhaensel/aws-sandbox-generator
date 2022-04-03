package models

import "github.com/graph-gophers/graphql-go"

type AWS string
type AZURE string
type GCP string

type Cloud struct {
	AWS   AWS
	AZURE AZURE
	GCP   GCP
}

var PublicCloud = Cloud{
	AWS:   "AWS",
	AZURE: "AZURE",
	GCP:   "GCP",
}

func (c Cloud) GetAWS() string {
	return string(c.AWS)
}
func (c Cloud) GetAZURE() string {
	return string(c.AZURE)
}

type LeaseSandBoxResult struct {
	Result interface{}
}

// SandboxMetadata and SandboxMetadataResolver
type SandboxMetadata struct {
	Id            graphql.ID
	AssignedUntil string
	AssignedSince string
	AssignedTo    string
}

/*
 SandboxMetadataResolver

 SandboxMetadata {
  id: ID!
  assignedUntil: String!
  assignedSince: String!
  assignedTo: String!
}
*/

type SandboxMetadataResolver struct{ cs *SandboxMetadata }

func (r *SandboxMetadataResolver) Id() graphql.ID {
	return r.cs.Id
}

func (r *SandboxMetadataResolver) AssignedUntil() string {
	return r.cs.AssignedUntil
}

func (r *SandboxMetadataResolver) AssignedSince() string {
	return r.cs.AssignedSince
}

func (r *SandboxMetadataResolver) AssignedTo() string {
	return r.cs.AssignedTo
}

// AwsSandbox and LeaseAwsResolver
type AwsSandbox struct {
	Metadata    SandboxMetadata
	AccountName string
}

type LeaseAwsResolver struct{ U AwsSandbox }

func (r *LeaseAwsResolver) Metadata() *SandboxMetadataResolver {
	return &SandboxMetadataResolver{&r.U.Metadata}
}

func (r *LeaseAwsResolver) AccountName() string {
	return r.U.AccountName
}

// AzureSandbox and LeaseAzureResolver

type AzureSandbox struct {
	Metadata    SandboxMetadata
	PipelineId  string
	Status      string
	ProjectId   string
	WebUrl      string
	SandboxName string
}

/*
 LeaseAzureResolver

 AzureSandbox {
  metadata: SandboxMetadata!
  sandboxName: String!
 }
*/
type LeaseAzureResolver struct{ U AzureSandbox }

func (r *LeaseAzureResolver) Metadata() *SandboxMetadataResolver {
	return &SandboxMetadataResolver{&r.U.Metadata}
}

func (r *LeaseAzureResolver) SandboxName() string {
	return r.U.SandboxName
}

// ToAwsSandbox and ToAzureSandbox

func (r *LeaseSandBoxResult) ToAwsSandbox() (*LeaseAwsResolver, bool) {
	c, ok := r.Result.(*LeaseAwsResolver)
	return c, ok
}

func (r *LeaseSandBoxResult) ToAzureSandbox() (*LeaseAzureResolver, bool) {
	c, ok := r.Result.(*LeaseAzureResolver)
	return c, ok
}
