package model

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/lambda"
	"log"
	"os"
)

var l *lambda.Lambda
var db *dynamodb.DynamoDB

func init() {
	if sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}); err != nil {
		log.Fatalf("Failed to connect to AWS: %s", err.Error())
	} else {
		db = dynamodb.New(sess)
		l = lambda.New(sess)
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
	out, _ := db.GetItem(&dynamodb.GetItemInput{TableName: m.Table(), Key: m.Key()})
	_ = dynamodbattribute.UnmarshalMap(out.Item, &m)
}

func Save(s Storable) {
	item, _ := dynamodbattribute.MarshalMap(s)
	_, err := db.PutItem(&dynamodb.PutItemInput{Item: item, TableName: s.Table()})
	fmt.Println(err)
}

func Delete(m Manageable) {
	_, _ = db.DeleteItem(&dynamodb.DeleteItemInput{Key: m.Key()})
}

func Call(f string, i ...interface{}) events.APIGatewayProxyResponse {
	r := events.APIGatewayProxyResponse{StatusCode: 500}
	var b []byte
	if i != nil && len(i) > 0 {
		b, _ = json.Marshal(&i[0])
	}
	if output, err := l.Invoke(&lambda.InvokeInput{FunctionName: aws.String(f), Payload: b}); err != nil {
		r.Body = err.Error()
	} else if err := json.Unmarshal(output.Payload, &r); err != nil {
		r.StatusCode = int(*output.StatusCode)
		r.Body = string(output.Payload)
	}
	return r
}
