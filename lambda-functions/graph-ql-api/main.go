package main

import (
	"context"
	"io/ioutil"
	"lambda/aws-sandbox/graph-ql-api/relay"
	"lambda/aws-sandbox/graph-ql-api/resolver"
	"lambda/aws-sandbox/graph-ql-api/schema"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	graphql "github.com/graph-gophers/graphql-go"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	relay := &relay.Handler{GraphqlSchema: graphqlSchema}

	return relay.ServeHTTP(ctx, request), nil
}

func init() {

}

func local() {
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {

		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		ctx := context.TODO()
		mock_handler := events.APIGatewayProxyRequest{
			Body: string(b),
		}
		responseJSON, _ := Handler(ctx, mock_handler)

		w.Write([]byte(responseJSON.Body))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	local()
	// lambda.Start(Handler)
}

// Test Sha, remove later
// Test Sha2, remove later :)
