package model

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"os"
)

var DB *dynamodb.DynamoDB

func init() {
	if sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}); err != nil {
		log.Fatalf("Failed to connect to AWS: %s", err.Error())
	} else {
		DB = dynamodb.New(sess)
	}
}

type Storable interface {
	Table() *string
}

type Findable interface {
	Storable
	Key() map[string]*dynamodb.AttributeValue
}

func Find(f Findable) {
	out, err := DB.GetItem(&dynamodb.GetItemInput{TableName: f.Table(), Key: f.Key()})
	if err != nil {
		fmt.Println(err)
	}

	if err = dynamodbattribute.UnmarshalMap(out.Item, &f); err != nil {
		fmt.Println(err)
	}
}
