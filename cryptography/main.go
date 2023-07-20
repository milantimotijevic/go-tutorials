package main

import (
	"crypto/sha256"
	"fmt"
)

var print = fmt.Println

func main() {
	print("-- Hashing --")

	someString := "Hello there!"

	hash := sha256.New()

	hash.Write([]byte(someString))

	hashedString := hash.Sum(nil)

	print(string(hashedString))
	fmt.Printf("%x\n", hashedString) // verb for converting it to hex (human readable)
}
