package api

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

// Required to transmit messages when CORS enabled in API Gateway.
var headers = map[string]string{"Access-Control-Allow-Origin": "*"}

func Log(r events.APIGatewayProxyRequest) {
	fmt.Printf("request: {\n"+
		"\tmethod: %s\n"+
		"\tresource: %s\n"+
		"\tpath: %s\n"+
		"\theaders: %v\n"+
		"\tquery_string_parameters: %v\n"+
		"\tbody: %s\n"+
		"\tbase64: %v\n"+
		"}\n",
		r.HTTPMethod, r.Resource, r.Path, r.Headers, r.QueryStringParameters, r.Body, r.IsBase64Encoded)
}

// Response returns an API Gateway Proxy Response with a nil error to provide detailed status codes and response bodies.
// While a status code must be provided, further arguments are recognized with reflection but not required.
func Response(code int, v ...interface{}) (events.APIGatewayProxyResponse, error) {
	var body string
	if v != nil && len(v) > 0 {
		b, _ := json.Marshal(v)
		body = string(b)
	}
	r := events.APIGatewayProxyResponse{StatusCode: code, Headers: headers, Body: body}
	fmt.Printf("response: {\n"+
		"\tcode: %d\n"+
		"\theaders: %v\n"+
		"\tbody: %s\n"+
		"}\n",
		r.StatusCode, r.Headers, r.Body)
	return r, nil
}
