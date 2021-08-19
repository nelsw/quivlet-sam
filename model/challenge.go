package model

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/nelsw/quivlet-sam/util/names"
	"github.com/nelsw/quivlet-sam/util/web"
	"reflect"
	"strconv"
)

type ChallengeResponse struct {
	Challenges []Challenge `json:"results"`
}

type Challenge struct {
	// Token is a unique 64 character alphanumeric string and identifier of a single Quivlet session.
	Token            string   `json:"token"`
	Index            int      `json:"index"`
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

func NewChallenge(token string) Challenge {
	r := &ChallengeResponse{}
	_ = json.Unmarshal(web.Get("https://opentdb.com/api.php?amount=1&type=multiple&encode=base64&difficulty=easy&token="+token), &r)
	return r.Challenges[0]
}

func (c *Challenge) SaveChallenge() {
	Save(c)
}

func (c *Challenge) FindChallenge() {
	Find(c)
}

func (c *Challenge) Table() *string {
	return aws.String(names.App + "_" + reflect.TypeOf(c).Elem().Name())
}

func (c *Challenge) Key() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"token": {S: aws.String(c.Token)},
		"index": {N: aws.String(strconv.Itoa(c.Index))},
	}
}
