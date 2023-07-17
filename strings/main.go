package main

import (
	"fmt"
)

func main() {
	fmt.Println("-- Go Strings --")
	var str = "Hello World"
	fmt.Println(str[0])   // byte value
	fmt.Println(str[0:1]) // slice it to get the "character" value

	for _, rune := range str { // even ranging over it initially produces "runes" (not proper "characters")
		var character string
		character += fmt.Sprintf("%c", rune)
		fmt.Println(character)
	}
}
