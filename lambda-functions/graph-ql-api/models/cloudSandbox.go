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
	Id            graphql.ID
	AssignedUntil string
	AssignedSince string
	AssignedTo    string
	AccountName   string
}

type LeaseAwsResolver struct {
	U AwsSandbox
}
type CloudSandbox interface {
	ID() graphql.ID
	Name() string
	Friends() *[]*characterResolver
	FriendsConnection(friendsConnectionArgs) (*friendsConnectionResolver, error)
	AppearsIn() []string
}
type human struct {
	ID   graphql.ID
	Name string
}
type humanResolver struct {
	h *human
}

func (r *humanResolver) ID() graphql.ID {
	return r.h.ID
}

func (r *humanResolver) Name() string {
	return r.h.Name
}

var humanData = make(map[graphql.ID]*human)

type CloudSandboxResolver struct {
	CloudSandbox
}
type character interface {
	ID() graphql.ID
	Name() string
	Friends() *[]*characterResolver
	AppearsIn() []string
}

type characterResolver struct {
	character
}
type droid struct {
	ID              graphql.ID
	Name            string
	PrimaryFunction string
}
type droidResolver struct {
	d *droid
}

func (r *droidResolver) ID() graphql.ID {
	return r.d.ID
}

func (r *droidResolver) Name() string {
	return r.d.Name
}
func (r *droidResolver) PrimaryFunction() *string {
	if r.d.PrimaryFunction == "" {
		return nil
	}
	return &r.d.PrimaryFunction
}

var droidData = make(map[graphql.ID]*droid)

func resolveCharacter(id graphql.ID) *characterResolver {
	if h, ok := humanData[id]; ok {
		return &characterResolver{&humanResolver{h}}
	}
	if d, ok := droidData[id]; ok {
		return &characterResolver{&droidResolver{d}}
	}
	return nil
}

func (r *LeaseAwsResolver) Id() graphql.ID {
	return r.U.Id
}

func (r *LeaseAwsResolver) AssignedUntil() string {
	return r.U.AssignedUntil
}

func (r *LeaseAwsResolver) AssignedSince() string {
	return r.U.AssignedSince
}

func (r *LeaseAwsResolver) AssignedTo() string {
	return r.U.AssignedTo
}

func (r *LeaseAwsResolver) AccountName() string {
	return r.U.AccountName
}

// AzureSandbox and LeaseAzureResolver

type AzureSandbox struct {
	Id            graphql.ID
	PipelineId    string
	AssignedUntil string
	AssignedSince string
	AssignedTo    string
	Status        string
	ProjectId     string
	WebUrl        string
}

type LeaseAzureResolver struct {
	U AzureSandbox
}

func (r *LeaseAzureResolver) Id() graphql.ID {
	return r.U.Id
}

func (r *LeaseAzureResolver) PipelineId() string {
	return r.U.PipelineId
}

func (r *LeaseAzureResolver) AssignedUntil() string {
	return r.U.AssignedUntil
}

func (r *LeaseAzureResolver) AssignedSince() string {
	return r.U.AssignedSince
}

func (r *LeaseAzureResolver) AssignedTo() string {
	return r.U.AssignedTo
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
