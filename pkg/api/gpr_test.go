package api

import (
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func TestLog(t *testing.T) {
	Log(events.APIGatewayProxyRequest{
		Resource:        "resource",
		Path:            "path",
		HTTPMethod:      "method",
		RequestContext:  events.APIGatewayProxyRequestContext{},
		Body:            "body",
		IsBase64Encoded: false,
	})
}

func TestResponse(t *testing.T) {

	c := 200
	r, _ := Response(c, &Problem{
		Category:         "Vehicles",
		Type:             "multiple",
		Difficulty:       "medium",
		Question:         "Which of the following passenger jets is the longest?",
		CorrectAnswer:    "Boeing 747-8",
		IncorrectAnswers: []string{"Boeing 747-8", "Airbus A330-200", "Boeing 747-8"},
	})
	if r.Headers["Access-Control-Allow-Origin"] != headers["Access-Control-Allow-Origin"] || r.StatusCode != c {
		t.Fail()
	}
}
