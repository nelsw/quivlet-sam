package model

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/nelsw/quivlet-sam/util/names"
	"github.com/nelsw/quivlet-sam/util/transform"
	"github.com/nelsw/quivlet-sam/util/web"
	"reflect"
	"time"
)

const tokenUrl = "https://opentdb.com/api_token.php?command=request"

// Session is an opentdb construct. They are described as unique keys that will help keep track of the questions we have
// received from said API. The Session Token guarantees that we will not receive the same question twice, until we have
// exhausted all ≈14,000 questions. Aside: Jeopardy episodes have 94 questions, incl. Daily Doubles and Final Jeopardy.
type Session struct {
	ID string `json:"id"`

	// Token is a unique 64 character alphanumeric string and identifier of a single Quivlet session.
	Token string `json:"token"`

	// Expiry is the Unix (millisecond) value, which defines the deadline for joining this session.
	Expiry int64 `json:"expiry"`
}

// NewSession sets the expiry of our future token first to hedge our time bets,
// retrieves a new token from an API call to the opentdb session token request endpoint,
// then sets said token on the session.
func (s *Session) NewSession() {
	transform.Unmarshal(web.Get(tokenUrl), s)
	s.ID = "id"
	s.Expiry = time.Now().Add(time.Hour * 6).Truncate(time.Minute).Unix()
}

func FindSession() *Session {
	s := new(Session)
	Find(s)
	return s
}

func (s *Session) SaveSession() {
	Save(s)
}

func (s *Session) DeleteSession() {
	Delete(s)
}

func (s *Session) Table() *string {
	return aws.String(names.App + "_" + reflect.TypeOf(s).Elem().Name())
}

func (s *Session) Key() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{"id": {S: aws.String("id")}}
}

func (s *Session) IsEmpty() bool {
	return s.Token == ""
}

func (s *Session) IsNotEmpty() bool {
	return !s.IsEmpty()
}

func (s *Session) IsExpired() bool {
	return s.Expiry == 0 || time.Now().After(time.Unix(s.Expiry, 0))
}

func (s *Session) IsNotExpired() bool {
	return !s.IsExpired()
}
