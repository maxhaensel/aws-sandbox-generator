package relay

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/graph-gophers/graphql-go"
)

type Handler struct {
	GraphqlSchema *graphql.Schema
}

type Params struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func (h *Handler) ServeHTTP(ctx context.Context, r events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	params := Params{}
	json.Unmarshal([]byte(r.Body), &params)

	response := h.GraphqlSchema.Exec(ctx, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal Server Error",
			StatusCode: 500,
			Headers:    map[string]string{},
		}
	}

	headers := map[string]string{
		"Access-Control-Allow-Origin": "*",
		"Content-Type":                "application/json",
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: 200,
		Headers:    headers,
	}
}
