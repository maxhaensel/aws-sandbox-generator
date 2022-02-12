package relay_test

import (
	"context"
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/relay"
	"lambda/aws-sandbox/graph-ql-api/resolver"
	"lambda/aws-sandbox/graph-ql-api/schema"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/graph-gophers/graphql-go"
)

var request events.APIGatewayProxyRequest

func TestServeHTTP(t *testing.T) {
	ctx := context.TODO()

	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	request = events.APIGatewayProxyRequest{
		Body: "",
	}

	relay := &relay.Handler{GraphqlSchema: graphqlSchema}

	result := relay.ServeHTTP(ctx, request)

	resultBody := result.Body == `{"errors":[{"message":"no operations in query document"}]}`
	statusCode := result.StatusCode == 200

	if !statusCode && !resultBody {
		t.Logf(fmt.Sprintf("expect statuscode 200 but got %v", result.StatusCode))
	}

}
