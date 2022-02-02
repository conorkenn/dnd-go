package dice

import (
	"math/rand"
	"time"
)

func Dice(max int) int {
	rand.Seed(time.Now().UnixNano())
	min := 1

	return rand.Intn(max-min+1) + min
}
