package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func TestHandleRequest(t *testing.T) {

	in := events.APIGatewayProxyRequest{
		Resource:                        "",
		Path:                            "",
		HTTPMethod:                      "",
		Headers:                         nil,
		MultiValueHeaders:               nil,
		QueryStringParameters:           map[string]string{"find": "all"},
		MultiValueQueryStringParameters: nil,
		PathParameters:                  nil,
		StageVariables:                  nil,
		RequestContext:                  events.APIGatewayProxyRequestContext{},
		Body:                            "",
		IsBase64Encoded:                 false,
	}
	out, _ := HandleRequest(in)
	fmt.Println(out)
}
