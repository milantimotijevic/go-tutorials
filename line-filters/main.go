package main

import (
	"bufio"
	"fmt"
	"os"
)

var print = fmt.Println

func main() {
	print("-- Line Filters --")

	scanner := bufio.NewScanner(os.Stdin)

	print("Say something")
	for scanner.Scan() {
		print("You said:", scanner.Text())
		print("Say something")
	}

	if err := scanner.Err(); err != nil {
		print("There's been an error with the scaner")
		print(err)
		os.Exit(1)
	}
}
