package model

import "testing"

func TestNewChallenge(t *testing.T) {
	p1 := NewChallenge(token.Value)
	p2 := NewChallenge(token.Value)
	if p1.Question == p2.Question {
		t.Fail()
	}
}
