package util

import "encoding/base64"

// FromBase64 is a convenience method for converting a Base64 encoded string to an ASCII string.
func FromBase64(s string) string {
	r, _ := base64.StdEncoding.DecodeString(s)
	return string(r)
}

// ToBase64 is a convenience method for encoding an ASCII string to Base64.
func ToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
