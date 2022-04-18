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

// AwsSandbox and LeaseAwsResolver
type AwsSandbox struct {
	Id            graphql.ID
	AssignedUntil string
	AssignedSince string
	AssignedTo    string
	AccountName   string
}

type LeaseAwsResolver struct {
	Aws AwsSandbox
}

func (r *LeaseAwsResolver) Id() graphql.ID {
	return r.Aws.Id
}

func (r *LeaseAwsResolver) AssignedUntil() string {
	return r.Aws.AssignedUntil
}

func (r *LeaseAwsResolver) AssignedSince() string {
	return r.Aws.AssignedSince
}

func (r *LeaseAwsResolver) AssignedTo() string {
	return r.Aws.AssignedTo
}

func (r *LeaseAwsResolver) AccountName() string {
	return r.Aws.AccountName
}

// AzureSandbox and LeaseAzureResolver

type AzureSandbox struct {
	Id            graphql.ID
	AssignedUntil string
	AssignedSince string
	AssignedTo    string
	Status        string
	ProjectId     string
	PipelineId    string
	WebUrl        string
}

type LeaseAzureResolver struct {
	Az AzureSandbox
}

func (r *LeaseAzureResolver) Id() graphql.ID {
	return r.Az.Id
}
func (r *LeaseAzureResolver) AssignedUntil() string {
	return r.Az.AssignedUntil
}
func (r *LeaseAzureResolver) AssignedSince() string {
	return r.Az.AssignedSince
}
func (r *LeaseAzureResolver) AssignedTo() string {
	return r.Az.AssignedTo
}
func (r *LeaseAzureResolver) Status() string {
	return r.Az.Status
}
func (r *LeaseAzureResolver) ProjectId() string {
	return r.Az.ProjectId
}
func (r *LeaseAzureResolver) PipelineId() string {
	return r.Az.PipelineId
}
func (r *LeaseAzureResolver) WebUrl() string {
	return r.Az.WebUrl
}

// CloudSandbox
type CloudSandbox interface {
	Id() graphql.ID
	AssignedUntil() string
	AssignedSince() string
	AssignedTo() string
}
type LeaseSandBoxResult struct {
	CloudSandbox
}
type Resolver struct{}

func (r *LeaseSandBoxResult) ToAwsSandbox() (*LeaseAwsResolver, bool) {
	c, ok := r.CloudSandbox.(*LeaseAwsResolver)
	return c, ok
}

func (r *LeaseSandBoxResult) ToAzureSandbox() (*LeaseAzureResolver, bool) {
	c, ok := r.CloudSandbox.(*LeaseAzureResolver)
	return c, ok
}
