package models

type ListSandboxeResponse struct {
	U ListSandboxeResolver
}

type ListSandboxeResolver struct {
	Sandboxes []*SandBoxItemResolver
}

type SandBoxResolver struct {
	U SandBoxResponse
}

type CreateNewSandboxResponse struct {
	U CreateNewSandboxResolver
}

type SandBoxResponse struct {
	Message string
	Sandbox SandBoxItemResolver
}

type CreateNewSandboxResolver struct {
	Sandbox SandBoxItemResolver
}

func (r *SandBoxResolver) Message() string {
	return r.U.Message
}
func (r *SandBoxResolver) Sandbox() *SandBoxItemResolver {
	return &r.U.Sandbox
}

func (r *CreateNewSandboxResponse) Sandbox() *SandBoxItemResolver {
	return &r.U.Sandbox
}

func (r *ListSandboxeResponse) Sandboxes() *[]*SandBoxItemResolver {
	return &r.U.Sandboxes
}

type SandboxItem struct {
	Account_id     string `dynamodbav:"account_id"`
	Account_name   string `dynamodbav:"account_name"`
	Assigned_since string `dynamodbav:"assigned_since"`
	Assigned_to    string `dynamodbav:"assigned_to"`
	Assigned_until string `dynamodbav:"assigned_until"`
	Available      string `dynamodbav:"available"`
}

type SandBoxItemResolver struct {
	U SandboxItem
}

func (r SandBoxItemResolver) Account_id() string {
	return r.U.Account_id
}
func (r *SandBoxItemResolver) Account_name() string {
	return r.U.Account_name
}
func (r *SandBoxItemResolver) Available() *string {
	return &r.U.Available
}
func (r *SandBoxItemResolver) Assigned_until() *string {
	return &r.U.Assigned_until
}
func (r *SandBoxItemResolver) Assigned_since() *string {
	return &r.U.Assigned_since
}
func (r *SandBoxItemResolver) Assigned_to() *string {
	return &r.U.Assigned_to
}
