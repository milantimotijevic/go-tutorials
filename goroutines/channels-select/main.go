package main

import (
	"fmt"
	"time"
)

func empower(startingValue int, increment int, delay int, channel chan int) {
	for {
		// need to use .Duration to set specific sleep duration
		time.Sleep(time.Duration(delay) * time.Millisecond)
		startingValue += increment
		channel <- startingValue
	}
}

func main() {
	channel1 := make(chan int)
	go empower(1, 1, 500, channel1)

	channel2 := make(chan int)
	go empower(1000, 1, 2000, channel2)

	for {
		// note, there would be a deadlock if one of the channels got closed
		// we would need a workaround for that
		select {
		// it seems I can use the same variable name for each case
		case msg1 := <-channel1:
			fmt.Println(msg1)
		case msg2 := <-channel2:
			fmt.Println(msg2)
		case <-time.After(4 * time.Second):
			// if nothing is received for 4 seconds, finish execution
			// time.After() also returns a channel
			return
		}
	}
}
