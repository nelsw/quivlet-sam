package model

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
	"github.com/nelsw/quivlet-sam/util/names"
	"reflect"
)

// User is the identity associated with a Quivlet contestant.
type User struct {
	*Token

	// ID is a 128 bit (16 byte) Universal Unique Identifier as defined in RFC 4122.
	ID uuid.UUID `json:"id"`

	// Name is a self defined or randomly generated username for friendly statistic reporting.
	Name string `json:"name"`

	// Nanos is a slice where each entry is equal to the amount of nanoseconds used to correctly answer a challenge.
	Nanos []int64 `json:"nanos"`
}

func (u *User) Table() *string {
	return aws.String(names.App + "_" + reflect.TypeOf(u).Elem().Name())
}
