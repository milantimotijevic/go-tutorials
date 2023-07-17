package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start:", time.Now())
	channel1 := make(chan string, 2)

	channel1 <- "Cat"
	channel1 <- "Dog"

	limiter := time.NewTicker(time.Millisecond * 500)
	// the channel is still open and not being fed, so a deadlock still happens in the sense
	// of idefinitely hanging, but it looks like the ticker prevents a panic from being thrown
	for channel1Item := range channel1 {
		<-limiter.C
		fmt.Println(channel1Item)
	}

	fmt.Println("End:", time.Now())
}
