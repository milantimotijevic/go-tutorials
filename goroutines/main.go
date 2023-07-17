package main

import (
	"fmt"
	"strconv"
	"time"
)

func transform(numbers []int64, channel chan string) {
	for _, number := range numbers {
		time.Sleep(time.Second)
		channel <- strconv.FormatInt(number, 10)
	}
	// need to close it because I'm consuming it with "range" in main
	close(channel)
}

func empower(numbers []int, channel chan int) {
	for _, number := range numbers {
		time.Sleep(time.Second)
		// no need to close it because I will iterate len() number of times in main
		channel <- number + 1
	}
}

func countDown(callType string, from int) {
	for i := from; i > 0; i-- {
		fmt.Println(callType, i)
	}
}

func main() {
	fmt.Println("-- Go Routines --")
	countDown("direct call one:", 3)
	countDown("direct call two:", 3)

	go countDown("routine call one:", 3)
	go countDown("routine call two:", 3)

	countDown("direct call three:", 3)

	time.Sleep(time.Second)

	fmt.Println("Examples with channels now - Transforming ints to strings")

	transformedNumsChan := make(chan string)
	go transform([]int64{1, 2, 3, 4, 5}, transformedNumsChan)

	for item := range transformedNumsChan {
		fmt.Println(item)
	}

	fmt.Println("Empowering now")
	empowerNumsChan := make(chan int)
	numsToEmpower := []int{1, 2, 3}
	go empower(numsToEmpower, empowerNumsChan)

	for i := 0; i < len(numsToEmpower); i++ {
		empowered := <-empowerNumsChan
		fmt.Println(empowered)
	}
}
