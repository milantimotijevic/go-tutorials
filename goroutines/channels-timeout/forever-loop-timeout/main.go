package main

import (
	"fmt"
	"time"
)

func extract(items []string, channel chan string, delay int) {
	for _, item := range items {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		channel <- fmt.Sprintf("%v@%v", item, delay)
	}
}

func main() {
	var names []string = []string{"Pera", "Mika", "Laza", "Zika"}
	channel1 := make(chan string)
	go extract(names, channel1, 500)

	channel2 := make(chan string)
	go extract(names, channel2, 2000)

	// on each iteration it will wait for one to execute, then switch to the next iteration
	// once all channels are exhausted, the iteration will naturally end up hitting the last case
	// and thus returning
	for {
		select {
		case nameCh1 := <-channel1:
			fmt.Println(nameCh1)
		case nameCh2 := <-channel2:
			fmt.Println(nameCh2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}
}
