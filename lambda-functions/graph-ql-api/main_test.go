package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	ctx := context.TODO()
	request := events.APIGatewayProxyRequest{
		Body: "",
	}

	result, err := Handler(ctx, request)

	resultErr := err == nil
	resultBody := result.Body == `{"errors":[{"message":"no operations in query document"}]}`
	statusCode := result.StatusCode == 200

	if !statusCode && !resultBody && !resultErr {
		t.Logf(fmt.Sprintf("expect statuscode 200 but got %v", result.StatusCode))
	}

}
