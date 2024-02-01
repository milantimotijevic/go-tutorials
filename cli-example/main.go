package main

import (
	"fmt"
	"os"
)

var print = fmt.Println

func main() {
	print("-- Cli Example --")
	print("Args Example:")
	// args are just an array
	args := os.Args
	print("All args:", args)

	print("Args excluding program path:", args[1:])
}
