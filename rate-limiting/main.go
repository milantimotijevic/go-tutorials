package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-- Go Rate Limiting --")

	requesters := []string{"Pera", "Mika", "Laza", "Zika", "Jeca", "Veca", "Neca"}
	// have to buffer it so I can fill it up in the main routine without simultaneously consuming
	// Alternatively, I could make it unbuffered and populate it in a subroutine
	requests := make(chan string, len(requesters))
	for _, requester := range requesters {
		requests <- requester
	}

	/*
		strangely enough, not closing it will not cause range to throw a deadlock, but rather
		hang indefinitely. Something to do with the limiter

		It is possible that the ticker somehow "refreshes" the loop, makes it somehow not be
		completely idle
	*/
	close(requests)

	limiter := time.NewTicker(time.Millisecond * 500)

	for requestItem := range requests {
		<-limiter.C
		fmt.Printf("Processed %v's request at %v\n", requestItem, time.Now())
	}

	fmt.Println("All done")
}
