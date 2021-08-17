package session

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/nelsw/quivlet-sam/pkg/api"
	"github.com/nelsw/quivlet-sam/pkg/util"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"
)

// Session is an opentdb construct. They are described as unique keys that will help keep track of the questions we have
// received from said API. The Session Token guarantees that we will not receive the same question twice, until we have
// exhausted all ≈14,000 questions. Aside: Jeopardy episodes have 94 questions, incl. Daily Doubles and Final Jeopardy.
type Session struct {

	// Token is the unique identifier of our session.
	Token string `json:"token"`

	// Expiry is the Unix (millisecond) value, which defines the expiration time of our session.
	Expiry int64 `json:"expiry"`
}

// saveNewToken sets the expiry of our future token first to hedge our time bets,
// retrieves a new token from an API call to the opentdb session token request endpoint,
// then sets said token on the session before finally persisting the new session to dynamo.
func (s *Session) saveNewToken() {

	s.Expiry = time.Now().Add(time.Hour * 6).Unix()

	request, err := http.NewRequest(http.MethodPost, api.OpenTDBTokenRequestUrl, nil)
	if err != nil {
		panic(err)
	}

	r := &api.TokenRequestResponse{}

	client := &http.Client{Timeout: 25 * time.Second}
	response, _ := client.Do(request)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(body, &r)
	_ = response.Body.Close()

	s.Token = r.Token

	item, _ := dynamodbattribute.MarshalMap(s)
	_, _ = db.PutItem(&dynamodb.PutItemInput{Item: item, TableName: table})
}

var (
	db    *dynamodb.DynamoDB
	key   = map[string]*dynamodb.AttributeValue{"id": {S: aws.String("key")}}
	table = aws.String(util.App + "_" + reflect.TypeOf(Session{}).String())
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

	// first try to retrieve a token from the db
	// while we can entertain multiple games in parallel,
	// we only use one Session token for a maximum of six hours.
	out, err := db.GetItem(&dynamodb.GetItemInput{TableName: table, Key: key})
	if err != nil {
		return api.Response(500, err)
	}

	// if we're here, the dynamo method completed as expected ...
	sess := &Session{}
	_ = dynamodbattribute.UnmarshalMap(out.Item, &sess)

	// but the result could be empty
	if sess == (&Session{}) {
		sess.saveNewToken()
	}

	// or it could be expired
	if time.Now().After(time.Unix(sess.Expiry, 0)) {
		_, _ = db.DeleteItem(&dynamodb.DeleteItemInput{Key: key})
		sess.saveNewToken()
	}

	return api.Response(200, &sess)
}
