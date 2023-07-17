package main

import (
	"fmt"
)

func main() {
	fmt.Println("-- Primitive Worker Example --")
	numChannel := make(chan int, 2)
	doneChannel := make(chan bool)

	go func() {
		for {
			// the loop will periodically get blocked based on how big a buffer we gave to numChannel
			item, open := <-numChannel
			if open {
				fmt.Println("Processed", item)
			} else {
				// if there are no more items to be read from numChannel, send a message to doneChannel
				// and terminate the loop
				fmt.Println("Pipeline complete")
				doneChannel <- true
				return
			}
		}
	}()

	numbers := []int{1, 2, 3, 4, 5}

	for _, number := range numbers {
		numChannel <- number
		fmt.Println("Sent", number)
	}

	close(numChannel)

	// forces the main routine to wait until doneChannel can be read from
	<-doneChannel
}
