package main

import (
	"fmt"
	"time"
)

func doStuffWithString(s string, channel chan string) {
	fmt.Println("Doing stuff...")
	time.Sleep(time.Second * 2)
	fmt.Println("Done")
	channel <- s + " changed"
}

func main() {
	fmt.Println("-- Channel Synchronization --")
	channel := make(chan string)
	go doStuffWithString("Hello", channel) // omitting the "go" keyword would cause a deadlock 

	// this is also considered consuming from a channel, even though I'm not using the value
	<-channel

	fmt.Println("Stuff with strings completed")
}
