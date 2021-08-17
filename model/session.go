package model

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/nelsw/quivlet-sam/util/names"
	"github.com/nelsw/quivlet-sam/util/transform"
	"github.com/nelsw/quivlet-sam/util/web"
	"time"
)

// Session is an opentdb construct. They are described as unique keys that will help keep track of the questions we have
// received from said API. The Session Token guarantees that we will not receive the same question twice, until we have
// exhausted all ≈14,000 questions. Aside: Jeopardy episodes have 94 questions, incl. Daily Doubles and Final Jeopardy.
type Session struct {
	ID string `json:"id"`

	*Token

	// Expiry is the Unix (millisecond) value, which defines the deadline for joining this session.
	Expiry int64 `json:"expiry,omitempty"`
}

// NewToken sets the expiry of our future token first to hedge our time bets,
// retrieves a new token from an API call to the opentdb session token request endpoint,
// then sets said token on the session.
func (s *Session) NewToken() {
	s.ID = "id"
	s.Expiry = time.Now().Add(time.Hour * 6).Unix()
	s.Token.New()
	transform.Unmarshal(web.Get(tokenUrl), s)
}

func FindToken() *Session {
	s := new(Session)
	Find(s)
	return s
}

func (s *Session) SaveToken() {
	Save(s)
}

func (s *Session) DeleteToken() {
	Delete(s)
}

func (s *Session) Table() *string {
	return aws.String(names.App + "_Session")
}

func (s *Session) Key() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{"id": {S: aws.String("id")}}
}

func (s *Session) IsEmpty() bool {
	return s.Token == nil
}

func (s *Session) IsNotEmpty() bool {
	return !s.IsEmpty()
}

func (s *Session) IsExpired() bool {
	return time.Now().After(time.Unix(s.Expiry, 0))
}

func (s *Session) IsNotExpired() bool {
	return !s.IsExpired()
}
