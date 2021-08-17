package transform

import "testing"

const (
	source = "R2VuZXJhbCBLbm93bGVkZ2U="
	target = "General Knowledge"
)

func TestToBase64(t *testing.T) {
	if ToBase64(target) != source {
		t.Fail()
	}
}

func TestFromBase64(t *testing.T) {
	if FromBase64(source) != target {
		t.Fail()
	}
}
