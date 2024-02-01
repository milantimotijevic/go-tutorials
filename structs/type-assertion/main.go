package main

import (
	"fmt"
)

type Person struct {
	specie string
	name   string
	age    int
}

func (p Person) introduce() {
	fmt.Printf("My name is %v, I am a(n) %v and I am %v years old", p.name, p.specie, p.age)
}

type Creature interface {
	introduce()
}

type Location interface {
	extractNPCs()
}

func main() {
	fmt.Println("-- Go Type Assertion --")
	// It's something like casting

	// assign it as an interface type, then we'll run some checks
	var pera Creature = Person{
		name: "Pera Peric",
		age:  33,
	}

	// check whether "pera" can be asserted as a "Person"
	// "peraPerson" will be the object of said type if it can
	// "peraIsPerson" will be true
	// otherwise, they will be nil and false respectively
	peraPerson, peraIsPerson := pera.(Person)
	fmt.Println(peraPerson, peraIsPerson)

	// since "pera" is not a Location, we will get nil and false respectively
	peraLocation, peraIsLocation := pera.(Location)
	fmt.Println(peraLocation, peraIsLocation)
}
