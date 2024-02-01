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

	// it will wait for exactly one channel to have a message in it, perform the associated expression
	// and exit the select block
	select {
	case s := <-stringChannel:
		fmt.Println("Received empowered string: ", s)
	case i := <-intChannel:
		fmt.Println("Received empowered int: ", i)
	}
}
