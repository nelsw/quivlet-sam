package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/nelsw/quivlet-sam/pkg/api"
	"github.com/nelsw/quivlet-sam/pkg/util"
	"log"
	"os"
	"reflect"
)

// User is the identity associated with a Quivlet contestant.
type User struct {

	// ID is a 128 bit (16 byte) Universal Unique Identifier as defined in RFC 4122.
	ID uuid.UUID `json:"id"`

	// Name is a self defined or randomly generated username for friendly statistic reporting.
	Name string `json:"name"`
}

var (
	db    *dynamodb.DynamoDB
	table = aws.String(util.App + "_" + reflect.TypeOf(User{}).String())
)

func init() {
	if sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}); err != nil {
		log.Fatalf("Failed to connect to AWS: %s", err.Error())
	} else {
		db = dynamodb.New(sess)
	}
}

func HandleRequest(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	api.Log(r)

	user := &User{}
	_ = json.Unmarshal([]byte(r.Body), &user)

	if user.ID.String() == "" {
		// this user is new, assign them an ID
		user.ID = uuid.New()
		// and add them to a slice of contents participants(?)
	}

	if user.Name == "" {
		// give them a moniker for round statistics
		user.Name = util.RandomName()
	}

	item, _ := dynamodbattribute.MarshalMap(&user)
	_, err := db.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: table,
	})

	if err != nil {
		return api.Response(500, err)
	}

	return api.Response(200)
}
