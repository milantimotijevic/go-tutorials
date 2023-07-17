package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-- Tickers --")
	ticker := time.NewTicker(time.Millisecond * 500)
	doneChannel := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-doneChannel:
				fmt.Println("Iterations complete")
				return // terminate the goroutine (even though it would get force-terminated upon the
				// exit from the main routine)
			case tick := <-ticker.C:
				fmt.Println("Tick at", tick)
			}
		}
	}()

	// the above routine will keep running while the main routine sleeps
	time.Sleep(time.Second * 2)
	// ticker gets stopped, which means it will no longer be an available option in select
	ticker.Stop()
	// done channel receives a value, which makes it a legit option for select
	doneChannel <- true
	fmt.Println("All done, exiting")
}
