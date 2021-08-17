package model

import (
	"github.com/nelsw/quivlet-sam/util/names"
	"testing"
	"time"
)

var s = &Session{}

func TestSession_Table(t *testing.T) {
	if *s.Table() != names.App+"_Session" {
		t.Fail()
	}
}

func TestSession_DeleteToken(t *testing.T) {
	s.DeleteToken()
}

func TestSession_NewToken(t *testing.T) {
	oldExpiry := s.Expiry
	s.NewToken()
	if oldExpiry == s.Expiry {
		t.Fail()
	}
}

func TestSession_SaveToken(t *testing.T) {
	oldExpiry := s.Expiry
	s.Expiry = time.Unix(oldExpiry, 0).Add(time.Minute).Unix()
	s.SaveToken()
	if oldExpiry == s.Expiry {
		t.Fail()
	}
}

func TestFindToken(t *testing.T) {
	oldValue := FindToken().Value
	s.NewToken()
	newValue := s.Value
	if oldValue == newValue {
		t.Fail()
	}
}

func TestSession_HelperMethods(t *testing.T) {
	s.IsEmpty()
	s.IsNotEmpty()
	s.IsExpired()
	s.IsNotExpired()
}
