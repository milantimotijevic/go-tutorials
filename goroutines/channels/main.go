package main

import (
	"fmt"
	"time"
)

func extract(names []string, channel chan string) {
	for _, name := range names {
		time.Sleep(1 * time.Second)
		channel <- name // writing to a channel is a blocking action
	}
	// important to close, or you'll end up with a deadlock
	close(channel)
}

func empowerNumbers(numbers []int, channel chan int) {
	for _, number := range numbers {
		time.Sleep(time.Second)
		// don't have to close it because I'll call it len() number of times in main
		channel <- number + 1
	}
}

func empowerSingle(number int, channel chan int) {
	time.Sleep(time.Second * 2)
	// no need to close since I'll read len() number of times
	channel <- number + 50
}

func main() {
	var names []string = []string{"Milan", "Milos", "Marko", "Pera", "Mika", "Laza", "Zika"}
	// no need to use a wait group, since channels can also orchestrate goroutines
	var channel chan string = make(chan string)
	go extract(names, channel)

	value, isOpen := <-channel // reading from a channel is also a blocking action
	fmt.Println(value)
	fmt.Println(isOpen)
	value, isOpen = <-channel
	fmt.Println(value)
	fmt.Println(isOpen)

	// iterate over the rest of the messages
	// channels behave like queues: once a message is read, it disappears
	for msg := range channel {
		fmt.Println(msg)
	}

	fmt.Println("Second Run")

	// need another one because the first one is closed, and that's permanent
	anotherChannel := make(chan string)
	go extract(names, anotherChannel)

	// this is a primitive way of looping over the messages in the channel
	for {
		name, isChannelOpen := <-anotherChannel
		if isChannelOpen {
			fmt.Println(name)
		} else {
			break
		}
	}

	fmt.Println("Empowering numbers now")
	numbersToEmpower := []int{1, 2, 3}
	empowerChannel := make(chan int)
	go empowerNumbers(numbersToEmpower, empowerChannel)

	for i := 0; i < len(numbersToEmpower); i++ {
		// shorthand for
		// empoweredNumber := <- empowerChannel
		fmt.Println(<-empowerChannel)
	}

	fmt.Println("Empowering single now (spawning threads inside for loop)")
	numbersToEmpowerSingle := []int{1, 2, 3}

	empowerSingleChannel := make(chan int)
	for _, number := range numbersToEmpowerSingle {
		go empowerSingle(number, empowerSingleChannel)
	}

	fmt.Println("Threads spawned")

	for i := 0; i < len(numbersToEmpowerSingle); i++ {
		empoweredNumber := <-empowerSingleChannel
		fmt.Println(empoweredNumber)
	}

	fmt.Println("Deadlock example")
	unfortunateChannel := make(chan string)
	unfortunateChannel <- "Hello"
	// I either need to make the channel buffered (with a capacity of at least 1)
	// or read from it in a different goroutine
	// otherwise the write will continuously wait for a read to happen, which never will
	// even if I place a read line in the main thread right below (write is waiting for a read
	// since it's unbuffered
	fmt.Println(<-unfortunateChannel) // this line does not help with an unbuffered channel
}
