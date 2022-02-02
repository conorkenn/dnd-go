package dice

import (
	"testing"
)

func TestDice(t *testing.T) {
	max := 12
	roll := Dice(max)

	if roll < 1 || roll >= max {
		t.Errorf("rolled a number we shouldn't have %d, expected a number between 1 and 6", roll)
	}
}
