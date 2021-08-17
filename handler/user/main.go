package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/nelsw/quivlet-sam/model"
	"github.com/nelsw/quivlet-sam/util/api"
	"github.com/nelsw/quivlet-sam/util/names"
)

func HandleRequest(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	api.Log(r)

	u := &model.User{}
	_ = json.Unmarshal([]byte(r.Body), &u)

	if u.ID.String() == "" { // this u is new, assign them an ID
		u.ID = uuid.New() // and add them to a slice of contents participants(?)
	}

	if u.Name == "" {
		u.Name = names.RandomName() // give them a moniker for round statistics
	}

	item, _ := dynamodbattribute.MarshalMap(&u)
	_, err := model.DB.PutItem(&dynamodb.PutItemInput{Item: item, TableName: u.Table()})

	if err != nil {
		return api.Response(500, err)
	}

	return api.Response(200)
}
