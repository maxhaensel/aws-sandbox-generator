package models

type SearchResult struct {
	Result interface{}
}

type AWS struct {
	Account_id string
}

type AWSResolver struct {
	U AWS
}

func (r *AWSResolver) Account_id() string {
	return r.U.Account_id
}

type AZURE struct {
	Pipeline_id string
}

type AZUREResolver struct {
	U AZURE
}

func (r *AZUREResolver) Pipeline_id() string {
	return r.U.Pipeline_id
}

func (r *SearchResult) ToAWS() (*AWSResolver, bool) {
	c, ok := r.Result.(*AWSResolver)
	return c, ok
}

func (r *SearchResult) ToAZURE() (*AZUREResolver, bool) {
	c, ok := r.Result.(*AZUREResolver)
	return c, ok
}
