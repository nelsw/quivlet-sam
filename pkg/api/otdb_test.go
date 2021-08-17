package api

import (
	"testing"
)

var token = NewToken()

func TestNewToken(t *testing.T) {
	if token == "" {
		t.Fail()
	}
}

func TestNewProblem(t *testing.T) {
	p1 := NewProblem(token)
	p2 := NewProblem(token)
	if p1.Question == p2.Question {
		t.Fail()
	}
}
