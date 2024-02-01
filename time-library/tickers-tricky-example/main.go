package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-- Tickers Tricky Example --")
	ticker := time.NewTicker(time.Millisecond * 500)
	doneChannel := make(chan bool, 1)
	counter := 0

	for {
		select {
		// the small 500 ms delay will actually cause the the first case to continuously
		// get blocked for a short period of time, which means on each iteration
		// the second case is also getting evaluated
		// It is only upon the second case's becoming available that the select statement
		// actually executes it
		case tick := <-ticker.C:
			fmt.Println("Tick", tick)
			counter--
			if counter <= 0 {
				doneChannel <- true
			}
		case <-doneChannel:
			fmt.Println("No more ticks")
			return
			//default:
			// this would get spammed like crazy in-between ticks
			//fmt.Println("Nothing")
		}
		// NOTE: it makes much more sense to place the "done" case first, to avoid evaluation spam
	}
}
