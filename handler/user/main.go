package main

import (
	"github.com/aws/aws-lambda-go/events"
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

	// if this user doesn't have a token, invite them to a session
	if u.Token == nil {
		out := model.Call("sessionHandler")
		var ss []model.Session
		transform.UnmarshalStr(out.Body, &ss)
		u.Token = ss[0].Token
		u.ID = uuid.New().String() // set a new uuid to create a composite key
	}

	// we MAY allow users to change their name during a session, if dev time permits.
	if u.Name == "" {
		u.Name = names.RandomName() // give them a moniker for round statistics
	}

	u.SaveUser()
	return api.Response(200, &u)
}
