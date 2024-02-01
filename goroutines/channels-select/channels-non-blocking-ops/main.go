package main

import (
	"fmt"
	"time"
)

func doSlowStuff(i int, channel chan int) {
	time.Sleep(time.Second)
	channel <- i
}

func main() {
	fmt.Println("-- Channels Non-Blocking Operations --")
	channel1 := make(chan int)
	channel2 := make(chan int)

	// it will evaluate them one by one and if neither is ready to be executed at the time of evaluation
	// it will move to evaluate the next one
	select {
	case v := <-channel1:
		fmt.Println("Received on channel1", v)
	case v := <-channel2:
		fmt.Println("Received on channel2", v)
	// since neither will be ready to be executed (both channels are empty), it will execute the default case
	default:
		fmt.Println("No activity")
	}
}
