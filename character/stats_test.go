package character

import (
	"testing"
)

func TestSetStats(t *testing.T) {
	maxRoll := 18
	statlist := setStats()

	if !(statlist.strength > 0 && statlist.strength <= maxRoll) {
		t.Errorf("Can't set stats higher than %d initially, tried to set %d", maxRoll, statlist.strength)
	}

	if !(statlist.dexterity > 0 && statlist.dexterity <= maxRoll) {
		t.Errorf("Can't set stats higher than %d initially, tried to set %d", maxRoll, statlist.dexterity)
	}

	if !(statlist.constitution > 0 && statlist.constitution <= maxRoll) {
		t.Errorf("Can't set stats higher than %d initially, tried to set %d", maxRoll, statlist.constitution)
	}

	if !(statlist.intellignece > 0 && statlist.intellignece <= maxRoll) {
		t.Errorf("Can't set stats higher than %d initially, tried to set %d", maxRoll, statlist.intellignece)
	}

	if !(statlist.wisdom > 0 && statlist.wisdom <= maxRoll) {
		t.Errorf("Can't set stats higher than %d initially, tried to set %d", maxRoll, statlist.wisdom)
	}

	if !(statlist.charisma > 0 && statlist.charisma <= maxRoll) {
		t.Errorf("Can't set stats higher than %d initially, tried to set %d", maxRoll, statlist.charisma)
	}
}
