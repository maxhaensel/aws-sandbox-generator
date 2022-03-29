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

// AwsSandbox and LeaseAwsResolver
type AwsSandbox struct {
	Id             graphql.ID
	Assigned_until string
	Assigned_since string
	Assigned_to    string
	Account_name   string
}

type LeaseAwsResolver struct {
	U AwsSandbox
}

func (r *LeaseAwsResolver) Id() graphql.ID {
	return r.U.Id
}

func (r *LeaseAwsResolver) Assigned_until() string {
	return r.U.Assigned_until
}

func (r *LeaseAwsResolver) Assigned_since() string {
	return r.U.Assigned_since
}

func (r *LeaseAwsResolver) Assigned_to() string {
	return r.U.Assigned_to
}

func (r *LeaseAwsResolver) Account_name() string {
	return r.U.Account_name
}

// AzureSandbox and LeaseAzureResolver

type AzureSandbox struct {
	Id             graphql.ID
	Pipeline_id    string
	Assigned_until string
	Assigned_since string
	Assigned_to    string
}

type LeaseAzureResolver struct {
	U AzureSandbox
}

func (r *LeaseAzureResolver) Id() graphql.ID {
	return r.U.Id
}

func (r *LeaseAzureResolver) Pipeline_id() string {
	return r.U.Pipeline_id
}

func (r *LeaseAzureResolver) Assigned_until() string {
	return r.U.Assigned_until
}

func (r *LeaseAzureResolver) Assigned_since() string {
	return r.U.Assigned_since
}

func (r *LeaseAzureResolver) Assigned_to() string {
	return r.U.Assigned_to
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
