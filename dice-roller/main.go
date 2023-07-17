package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomNumber(min int, max int) int {
	randomValue := rand.Intn(max-min+1) + min

	return randomValue
}

type Ability struct {
	name      string
	diceCount int
	diceSides int
}

func cast(ability Ability) int {
	var castValue int = 0
	for i := 0; i < ability.diceCount; i++ {
		castValue += getRandomNumber(1, ability.diceSides)
	}

	return castValue
}

func main() {
	const TURNS_COUNT int = 50000
	rand.Seed(time.Now().UnixNano())

	thunderwave := Ability{
		name:      "Thunderwave",
		diceCount: 2,
		diceSides: 8,
	}

	thunderwaveTotal := 0

	for i := 0; i <= TURNS_COUNT; i++ {
		thunderwaveTotal += cast(thunderwave)
	}

	fmt.Printf("\n Thunderwave:\n Total: %v\n Turns: %v\n Average: %v\n", thunderwaveTotal, TURNS_COUNT, thunderwaveTotal/TURNS_COUNT)

	burningHands := Ability{
		name:      "Burning Hands",
		diceCount: 3,
		diceSides: 6,
	}

	burningHandsTotal := 0

	for i := 0; i <= TURNS_COUNT; i++ {
		burningHandsTotal += cast(burningHands)
	}

	fmt.Printf("\n Burning Hands:\n Total: %v\n Turns: %v\n Average: %v\n", burningHandsTotal, TURNS_COUNT, burningHandsTotal/TURNS_COUNT)
}
