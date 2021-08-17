package model

import (
	"testing"
)

var token = Token{}

func init() {
	token.New()
}

func TestNewToken(t *testing.T) {
	if token.Value == "" {
		t.Fail()
	}
	if NewToken(token.Value).Value != token.Value {
		t.Fail()
	}
}
