package main

import (
	"flag"
	"fmt"
)

var print = fmt.Println

func main() {
	print("-- Flags Example --")

	isDarkUrge := false

	namePtr := flag.String("name", "Tav", "Character Name")
	classPtr := flag.String("class", "", "Playable Class (e.g. Wizard)")
	subclassPtr := flag.String("subclass", "", "Playable Character Subclass")
	level := flag.Int("level", 1, "Character level")
	// can also attach it to an existing var
	flag.BoolVar(&isDarkUrge, "isDarkUrge", false, "Origin")

	flag.Parse()

	print("Selected name:", *namePtr)
	print("Selected class:", *classPtr)
	print("Selected subclass:", *subclassPtr)
	print("Selected level:", *level)
	print("Is dark urge:", isDarkUrge)
}
