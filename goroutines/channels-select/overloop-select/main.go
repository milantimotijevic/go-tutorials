package main

import (
	"fmt"
	"time"
)

func empowerString(s string, channel chan string) {
	fmt.Println("Empowering: ", s)
	time.Sleep(time.Second * 3)
	channel <- "super " + s
}

func empowerInt(i int, channel chan int) {
	fmt.Println("Empowering: ", i)
	time.Sleep(time.Second * 3)
	channel <- i + 1
}

func main() {
	stringChannel := make(chan string)
	intChannel := make(chan int)

	go empowerString("cat", stringChannel)
	go empowerInt(1, intChannel)

	// this will cause a deadlock because we're looping 3 times but only have 2 writes
	// across all channels
	// the deadlock occurres because the main routine would get stuck waiting for something
	// that can never happen
	for i := 0; i < 3; i++ {
		select {
		case s := <-stringChannel:
			fmt.Println("Received empowered string: ", s)
		case i := <-intChannel:
			fmt.Println("Received empowered int: ", i)
		}
	}
}
