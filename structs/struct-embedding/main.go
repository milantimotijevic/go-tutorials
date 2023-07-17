package main

import (
	"fmt"
)

type GameObject interface {
	identify()
}

type Creature struct {
	specie string
	health int
}

func (c Creature) identify() {
	fmt.Printf("I am a(n) %v and my health is %v\n", c.specie, c.health)
}

type Character struct {
	name string
	Creature
}

func (char Character) introduce() {
	fmt.Printf("My name is %v\n", char.name)
}

func main() {
	fmt.Println("-- Go Structs Embedding --")

	humanCreature := Creature{
		specie: "Human",
		health: 10,
	}

	// Creature implements GameObject, so it can be stored under its reference
	var elfCreature GameObject = Creature{
		specie: "Elf",
		health: 9,
	}

	elfCreature.identify()

	// Creature implements GameObject and Character has an embedded Creature
	// so a Character also implements GameObject
	var milanChar GameObject = Character{
		name:     "Milan",
		Creature: humanCreature,
	}

	milanChar.identify()

}
