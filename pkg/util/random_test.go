package util

import "testing"

func TestShuffle(t *testing.T) {
	temp := names[0:99]
	n1 := temp[0]
	Shuffle(temp)
	n2 := temp[0]
	if n1 == n2 {
		t.Fail()
	}
}

func TestRandomInt(t *testing.T) {
	min := 0
	max := 999
	n1 := RandomInt(min, max)
	n2 := RandomInt(min, max)
	if n1 == n2 {
		t.Fail()
	}
}
