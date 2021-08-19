package model

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/nelsw/quivlet-sam/util/names"
	"reflect"
)

// User is the identity associated with a Quivlet contestant.
type User struct {

	// Token is a unique 64 character alphanumeric string and identifier of a single Quivlet session.
	Token string `json:"token"`

	// ID is the string representation of a 128 bit (16 byte) Universal Unique Identifier as defined in RFC 4122.
	ID string `json:"id"`

	// Name is a self defined or randomly generated username for friendly statistic reporting.
	Name string `json:"name"`

	// Result is a slice of int64's where each value represents a correctly answered challenge in a certain time period.
	Result []int64 `json:"result"`

	// Eliminated is used to flag a user that answered a challenge incorrectly and no longer part of the contest.
	Eliminated bool `json:"eliminated"`

	Save bool
}

func (u *User) Table() *string {
	return aws.String(names.App + "_" + reflect.TypeOf(u).Elem().Name())
}

func (u *User) SaveUser() {
	SaveUser(u)
}
