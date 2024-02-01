package main

import (
	"fmt"
)

type Character struct {
	name      string
	health    int
	abilities []Ability
}

var player = Character{
	name:      "Mitlandir",
	health:    8,
	abilities: []Ability{fireBolt, rayOfFrost},
}

var enemy = Character{
	name:      "Hobgoblin",
	health:    15,
	abilities: []Ability{slam},
}

func announceStats() {
	fmt.Println("#############")
	fmt.Printf("%v's health: %v\n", player.name, player.health)
	fmt.Printf("%v's health: %v\n", enemy.name, enemy.health)
	fmt.Println("#############")
}
