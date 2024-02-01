package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("-- Wait Group Closure Example --")
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		/*
			We are assigning a new i value. Otherwise, by the time each closure spins up, the loop
			will have reached i === 5, and they would all capture that value

			Alternatively, we could make it so the closure function takes in an int argument, pass i to it
			and use that one inside the closure
		*/
		i := i
		go func() {
			defer wg.Done()
			fmt.Println("Worker", i, "starting")
			time.Sleep(time.Second)
			fmt.Println("Worker", i, "finished")
		}()
	}

	wg.Wait()
	fmt.Println("All done")
}
