package main

import (
	"fmt"
)

type Creature struct {
	name   string
	specie string
	age    int
}

// methods are created using this weird syntax
// it's called a "receiver type"
func (c Creature) introduceOneself() {
	fmt.Printf("My name is %v, I am a(n) %v and I am %v year(s) old\n", c.name, c.specie, c.age)
}

// can also attach a method using a pointer, probably to save memory and to allow mutation
func (c *Creature) introduceOneselfLoudly() {
	fmt.Printf("MY NAME IS %v, I AM A(N) %v AND I AM %v YEAR(S) OLD! AARRRRRRR!\n", c.name, c.specie, c.age)
}

func main() {
	fmt.Println("-- Go Structs --")

	// anonymous struct declaration
	dog := struct {
		name  string
		breed string
	}{
		name:  "Woofer",
		breed: "Good Boy",
	}

	fmt.Println(dog)

	person := Creature{
		name:   "Milan",
		specie: "Human",
		age:    33,
	}

	fmt.Println(person)

	person.introduceOneself()
	person.introduceOneselfLoudly()
}
