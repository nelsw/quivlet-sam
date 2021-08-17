package random

import (
	"math/rand"
	"time"
)

// init is responsible for seeding random numbers package.
// Without a unique seed, our RandomName method will return
// the same name in perpetuity.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Shuffle rearranges given slice strings in a random order.
func Shuffle(slice []string) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// RandomInt returns a random int between the given bounds.
func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
