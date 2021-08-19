package model

import "testing"

func TestNewChallenge(t *testing.T) {
	s := Session{}
	s.NewSession()
	p1 := NewChallenge(s.Token)
	p2 := NewChallenge(s.Token)
	if p1.Question == p2.Question {
		t.Fail()
	}
}
