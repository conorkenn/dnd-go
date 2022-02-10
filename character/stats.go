package character

import (
	"dnd/dice"
)

type statlist struct {
	strength     int
	dexterity    int
	constitution int
	intellignece int
	wisdom       int
	charisma     int
}

func setStats() *statlist {
	s := statlist{
		strength:     dice.Dice(18),
		dexterity:    dice.Dice(18),
		constitution: dice.Dice(18),
		intellignece: dice.Dice(18),
		wisdom:       dice.Dice(18),
		charisma:     dice.Dice(18),
	}

	return &s
}
