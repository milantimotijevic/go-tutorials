package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-- Sending In Select --")

	channel1 := make(chan string)
	channel2 := make(chan int)

	go func() {
		fmt.Println("Read from subroutine", <-channel2)
	}()

	// ensure above goroutine has enough time to start before we reach the select statement
	time.Sleep(time.Millisecond * 100)

	// it will evaluate the first case, figure out the channel is empty and move onto evaluating the next
	select {
	case value := <-channel1:
		fmt.Println("Received from channel1", value)
	// since we've given it 100 ms, by the time this case is evalualted the subroutine will have started
	case channel2 <- 5: // we can send stuff inside a select block
		fmt.Println("Send to channel")
	default:
		fmt.Println("No activity")
	}
}
