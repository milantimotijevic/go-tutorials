package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-- Bursty Limiter --")

	requesters := []string{"Pera", "Mika", "Laza", "Zika", "Jeca", "Veca", "Neca"}
	requests := make(chan string, len(requesters))
	for _, requester := range requesters {
		requests <- requester
	}
	close(requests)

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		// add exactly 3 time objects into this channel
		// these three will be immediately consumable later on, without blocking
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 500) {
			// continuously add a new time object into the channel every 500 ms
			// an item will only be added once at least one has been consumed
			burstyLimiter <- t
		}
	}()

	for requestItem := range requests {
		// since it already has 3 items ready to be consumed, there will be no blocking on the
		// initial three iterations
		<-burstyLimiter
		fmt.Printf("Processed %v's request at %v\n", requestItem, time.Now())
	}

	fmt.Println("All done")
}
