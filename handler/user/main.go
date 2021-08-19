package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	"github.com/nelsw/quivlet-sam/model"
	"github.com/nelsw/quivlet-sam/util/api"
	"github.com/nelsw/quivlet-sam/util/names"
	"github.com/nelsw/quivlet-sam/util/transform"
)

// HandleRequest is responsible for
// 1. associating a model.User with a model.Session
// 2. saving the nanosecond duration required to solve a model.Challenge
func HandleRequest(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	api.Log(r)

	u := &model.User{}
	transform.UnmarshalStr(r.Body, &u)

	if r.QueryStringParameters["find"] == "all" {
		users := model.FindUsers(u.Token)
		return api.Response(200, &users)
	}

	if !u.Save {
		return api.Response(200)
	}

	if r.QueryStringParameters["new"] == "id" {
		out := model.Call("sessionHandler")
		var s model.Session
		transform.UnmarshalStr(out.Body, &s)
		u.Token = s.Token
		u.ID = uuid.New().String() // set a new uuid to create a composite key
		if u.Name == "" {
			u.Name = names.RandomName() // give them a moniker for round statistics
		}
		model.SaveUser(u)
		return api.Response(200, &u)
	}

	// if we're here, we're either a new user, or simply persisting user data
	if r.QueryStringParameters["save"] == "id" {
		model.SaveUser(u)
		return api.Response(200, &u)
	}

	return api.Response(400)
}

func main() {
	lambda.Start(HandleRequest)
}
