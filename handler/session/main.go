package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nelsw/quivlet-sam/model"
	"github.com/nelsw/quivlet-sam/util/api"
	"time"
)

func HandleRequest(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	api.Log(r)

	s := model.FindSession()

	if r.QueryStringParameters["find"] == "id" {
		//return api.Response(200, &s)
	}

	if s.IsExpired() || r.QueryStringParameters["refresh"] == "id" {
		s.DeleteSession()
		s.NewSession()
		s.SaveSession()
		return api.Response(200, &s)
	}

	sessionSize := len(*model.FindUsers(s.Token))

	if sessionSize > 1 {
		// give future users ≈2 more minutes to join
		s.Expiry = time.Now().Add(time.Minute * 1).Truncate(time.Minute).Unix()
		s.SaveSession()
		return api.Response(200, &s)
	}

	if sessionSize > 0 {
		// PUT a new challenge in the db to initiate the contest
		_ = model.Call("challengeHandler", &model.Challenge{Token: s.Token, Index: 0})
		return api.Response(200, &s)
	}

	return api.Response(200, &s)
}

func main() {
	lambda.Start(HandleRequest)
}
