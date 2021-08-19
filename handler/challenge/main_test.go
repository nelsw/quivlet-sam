package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/nelsw/quivlet-sam/model"
	"github.com/nelsw/quivlet-sam/util/transform"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	c := model.Challenge{
		Token: model.FindSession().Token,
		Index: 0,
	}
	b := transform.Marshal(&c)
	s := string(b)
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
		Body:                            s,
		IsBase64Encoded:                 false,
	}
	out, _ := HandleRequest(in)
	fmt.Println(out)
}
