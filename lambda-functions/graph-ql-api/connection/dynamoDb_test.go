package connection_test

import (
	"context"
	"lambda/aws-sandbox/graph-ql-api/connection"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func TestGetDynamoDbClient(t *testing.T) {
	svc := connection.GetDynamoDbClient(context.TODO())

	result := reflect.TypeOf(*svc) == reflect.TypeOf(dynamodb.Client{})

	if result == false {
		t.Errorf("error in item")
	}
}
