package main

import (
	"fmt"
)

type RecreationalItem interface {
	// every struct that has this method is considered to implement this interface
	printInfo()
}

type BlankInterface interface {
	// every single thing technically implements this
}

type Game struct {
	name  string
	price float32
}

// "class methods" are assigned using this weird syntax
func (g Game) printInfo() {
	fmt.Printf("Game %v costs $%v\n", g.name, g.price)
}

type Book struct {
	name  string
	price float32
}

func (b Book) printInfo() {
	fmt.Printf("Book %v costs $%v\n", b.name, b.price)
}

func main() {
	aom := Game{
		name:  "Age of Mythology",
		price: 15,
	}

	cof := Book{
		name:  "Children of Fire",
		price: 12,
	}

	aom.printInfo()
	cof.printInfo()

	var deusEx RecreationalItem = Game{
		name:  "Deus Ex Machina",
		price: 7,
	}

	deusEx.printInfo()

	var bg3 BlankInterface = Game{ // implements this blank interface because everything does
		name:  "Baldur's Gate 3",
		price: 70,
	}

	fmt.Println(bg3)

	// can quite literally stuff anything I want into a slice if I consider it a member of a blank interface
	var sliceOfStuff []BlankInterface = []BlankInterface{"cat", bg3, aom, 15}
	for index, item := range sliceOfStuff {
		fmt.Printf("Index %v item %v\n", index, item)
	}

	wow := Game{"World of Warcraft", 55} // strange syntax for a constructor, but oh well
	fmt.Println(wow)
	// TODO look into generics
}
