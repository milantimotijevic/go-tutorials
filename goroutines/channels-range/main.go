package main

import (
	"fmt"
)

func main() {
	fmt.Println("-- Channels Range --")
	// have to make it buffered so I can fill it up without blocking the main routine
	// OR, I could make it unbuffered and populate it in a subroutine
	numChannel := make(chan int, 5)

	numbers := []int{1, 2, 3, 4}

	for _, number := range numbers {
		numChannel <- number
	}

	// it's important to close it, otherwise below range statement will cause a deadlock
	close(numChannel)

	// can read from a closed channel
	fmt.Println(<-numChannel)
	fmt.Println("Range:")
	// range will pick up where the previous code left off
	for item := range numChannel {
		// this won't end up blocking because numChannel has been closed
		// leaving it open would cause "range" to block the main routine
		fmt.Println(item)
	}
}
