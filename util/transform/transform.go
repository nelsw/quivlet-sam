package transform

import (
	"encoding/base64"
	"encoding/json"
)

// FromBase64 is a convenience method for converting a Base64 encoded string to an ASCII string.
func FromBase64(s string) string {
	r, _ := base64.StdEncoding.DecodeString(s)
	return string(r)
}

// ToBase64 is a convenience method for encoding an ASCII string to Base64.
func ToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Unmarshal is a convenience method for unmarshalling a byte array into the given struct.
func Unmarshal(b []byte, v interface{}) {
	_ = json.Unmarshal(b, v)
}

// UnmarshalStr is a convenience method for unmarshalling a string into the given struct.
func UnmarshalStr(s string, v interface{}) {
	_ = json.Unmarshal([]byte(s), v)
}
