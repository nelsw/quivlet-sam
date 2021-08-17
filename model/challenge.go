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

type ChallengeKey struct {
	*Token
	*Index
}

type Challenge struct {
	*ChallengeKey
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

// Index identifies the challenge number given to Quivlet participants, beginning with 1.
type Index struct {
	Value int `json:"index"`
}

func NewIndex(value int) *Index {
	index := new(Index)
	index.Value = value
	return index
}

func NewChallenge(token string) Challenge {
	r := &ChallengeResponse{}
	_ = json.Unmarshal(web.Get("https://opentdb.com/api.php?amount=1&type=multiple&encode=base64&token="+token), &r)
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
		"token": {S: aws.String(c.Token.Value)},
		"index": {N: aws.String(strconv.Itoa(c.Index.Value))},
	}
}
