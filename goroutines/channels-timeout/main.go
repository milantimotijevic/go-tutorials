package main

import (
	"fmt"
	"time"
)

func doSlowStuff(s string, channel chan string, timeout int) {
	time.Sleep(time.Second * time.Duration(timeout))
	channel <- s
}

func main() {
	fmt.Println("-- Channels Timeout --")
	channel1 := make(chan string)

	go doSlowStuff("hi", channel1, 2)

	select {
	case v := <-channel1:
		fmt.Println("Read from channel1", v)
	case <-time.After(time.Second * 3): // if nothing else gets read, this statement will
		fmt.Println("Timeout")
	}
}
