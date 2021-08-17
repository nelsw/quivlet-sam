package model

import (
	"encoding/json"
	"github.com/nelsw/quivlet-sam/util/web"
)

const tokenUrl = "https://opentdb.com/api_token.php?command=request"

// Token is a unique 64 character alphanumeric string and identifier of a single Quivlet session.
type Token struct {
	Value string `json:"token"`
}

func NewToken(value string) *Token {
	token := new(Token)
	token.Value = value
	return token
}

func (t *Token) New() {
	_ = json.Unmarshal(web.Get(tokenUrl), t)
}
