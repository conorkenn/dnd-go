package character

type character struct {
	name  string
	stats statlist
}

func CreateCharacter() *character {
	c := character{
		name:  "billy",
		stats: *setStats(),
	}
	return &c
}
