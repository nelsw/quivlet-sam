package names

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	name1 := RandomName()
	name2 := RandomName()
	if name1 == name2 {
		t.Fail()
	}
	fmt.Println(time.Now().Truncate(time.Minute))
}
