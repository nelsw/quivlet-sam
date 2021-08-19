package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/nelsw/quivlet-sam/model"
	"testing"
)

func TestHandleRequest_LiveSession(t *testing.T) {

	in := events.APIGatewayProxyRequest{
		Resource:                        "",
		Path:                            "",
		HTTPMethod:                      "",
		Headers:                         nil,
		MultiValueHeaders:               nil,
		QueryStringParameters:           nil,
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

func TestHandleRequest_NoSession(t *testing.T) {
	s := model.FindSession()
	s.DeleteSession()
	in := events.APIGatewayProxyRequest{
		Resource:                        "",
		Path:                            "",
		HTTPMethod:                      "",
		Headers:                         nil,
		MultiValueHeaders:               nil,
		QueryStringParameters:           nil,
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
