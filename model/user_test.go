package model

import (
	"github.com/nelsw/quivlet-sam/util/names"
	"testing"
)

var u = &User{}

func TestUser_Table(t *testing.T) {
	if *u.Table() != names.App+"_User" {
		t.Fail()
	}
}
