package main

import (
	"dnd/character"
	"fmt"
)

func main() {
	myBoi := character.CreateCharacter()

	fmt.Print(myBoi)
}
