package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/nelsw/quivlet-sam/model"
	"github.com/nelsw/quivlet-sam/util/api"
	"time"
)

func HandleRequest(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	api.Log(r)

	s := &model.Session{}

	// first try to retrieve a token from the db
	// while we can entertain multiple games in parallel,
	// we only use one Session token for 14K questions or
	// six hours of inactivity ... whichever comes first.
	out, _ := model.DB.GetItem(&dynamodb.GetItemInput{TableName: s.Table(), Key: s.Key()})

	// if we're here, the dynamo method completed as expected ...
	_ = dynamodbattribute.UnmarshalMap(out.Item, &s)

	if s.IsNotEmpty() && s.IsNotExpired() {
		// todo - create the first challenge if it doesn't already exist

		// if we're here, we have at least 2 users who have joined the session
		// let's give potential future participants another minute to join
		s.Expiry = time.Now().Add(time.Hour).Unix()

	} else if s.IsEmpty() {
		fmt.Println("empty")
		// only occurs on system initialization
		s.SaveNewToken()
	} else if s.IsExpired() {
		fmt.Println("expired")
		// this user missed the session deadline
		_, _ = model.DB.DeleteItem(&dynamodb.DeleteItemInput{Key: s.Key()})
		s.SaveNewToken()
	}

	return api.Response(200, &s)
}

func main() {
	lambda.Start(HandleRequest)
}
