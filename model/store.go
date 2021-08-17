package model

import (
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

type Manageable interface {
	Storable
	Key() map[string]*dynamodb.AttributeValue
}

func Find(m Manageable) {
	out, _ := DB.GetItem(&dynamodb.GetItemInput{TableName: m.Table(), Key: m.Key()})
	_ = dynamodbattribute.UnmarshalMap(out.Item, &m)
}

func Save(s Storable) {
	item, _ := dynamodbattribute.MarshalMap(s)
	_, _ = DB.PutItem(&dynamodb.PutItemInput{Item: item, TableName: s.Table()})
}

func Delete(m Manageable) {
	_, _ = DB.DeleteItem(&dynamodb.DeleteItemInput{Key: m.Key()})
}
