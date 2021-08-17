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

	s := model.FindToken()

	// if the session is active, we must have at least 2 users who have joined the session
	if s.IsNotEmpty() && s.IsNotExpired() {
		// prepare the first challenge
		ck := &model.ChallengeKey{s.Token, model.NewIndex(0)}
		_ = model.Call("challengeHandler", &ck)
		// let's give potential future participants another minute to join
		s.Expiry = time.Now().Add(time.Minute).Unix()
		return api.Response(200, &s)
	}

	// if the session table is empty ...
	if s.IsEmpty() {
		s.NewToken()
		s.SaveToken()
		return api.Response(200, &s)
	}

	// else this user missed the session deadline
	s.DeleteToken()
	s.NewToken()
	s.SaveToken()
	return api.Response(200, &s)
}

func main() {
	lambda.Start(HandleRequest)
}
