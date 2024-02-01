package main

import (
	"fmt"
)

func main() {
	fmt.Println("-- Go Maps --")

	var companions map[string]string = make(map[string]string)
	companions["Astarion"] = "Rogue"
	companions["Gale"] = "Wizard"
	companions["Shadowheart"] = "Cleric"
	companions["Wyll"] = "Warlock"
	companions["Lae'zel"] = "Fighter"
	fmt.Println(companions)

	delete(companions, "Lae'zel")
	fmt.Println(companions)
	fmt.Println(companions["Lae'zel"])

	for k, v := range companions {
		fmt.Printf("Key %v, Value %v\n", k, v)
	}
}
