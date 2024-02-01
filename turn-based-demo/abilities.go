package main

import (
	"fmt"
)

type Ability struct {
	name           string
	minValue       int
	maxValue       int
	getCastMessage func(casterName string, targetName string, value int) string
}

var fireBolt = Ability{
	name:     "Fire Bolt",
	minValue: 0,
	maxValue: 10,
	getCastMessage: func(casterName string, targetName string, abilityValue int) string {
		return fmt.Sprintf("%v casts Fire Bolt at %v for %v damage\n", casterName, targetName, abilityValue)
	},
}

var rayOfFrost = Ability{
	name:     "Ray of Frost",
	minValue: 0,
	maxValue: 8,
	getCastMessage: func(casterName string, targetName string, abilityValue int) string {
		return fmt.Sprintf("%v casts Ray of Frost at %v for %v damage\n", casterName, targetName, abilityValue)
	},
}

var slam = Ability{
	name:     "Slam",
	minValue: 0,
	maxValue: 6,
	getCastMessage: func(casterName string, targetName string, abilityValue int) string {
		return fmt.Sprintf("%v slams %v for %v damage\n", casterName, targetName, abilityValue)
	},
}
