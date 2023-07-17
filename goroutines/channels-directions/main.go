package main

import (
	"fmt"
)

// the direction is function-specific
// it does not change the channel itself, only what the function (scope?) can do with it
func enqueueString(queueChan chan<- string, str string) {
	queueChan <- str
}

func superifyString(queueChan <-chan string, resultChan chan<- string) {
	item := <-queueChan
	fmt.Printf("Superifying %v\n", item)
	item = "SUPER " + item
	resultChan <- item
}

func main() {
	fmt.Println("-- Channels Directions --")

	queueChan := make(chan string)
	// In this example, I don't have to spawn a routine for enqueueString and superifyString
	// but I would need to buffer the channels in that case, otherwise we would get a deadlock
	go enqueueString(queueChan, "Saiyan")

	resultChan := make(chan string)
	go superifyString(queueChan, resultChan)

	// this line is the only reason the main routine ends up waiting for the other two routines
	fmt.Println(<-resultChan)
}
