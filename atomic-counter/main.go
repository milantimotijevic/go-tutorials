package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("-- Go Atomic Counter --")

	var counter1 uint64
	var wg1 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg1.Add(1)
		go func() {
			//counter1 += 1 // unreliable, because goroutines would interfere among each other
			atomic.AddUint64(&counter1, 1) // consistent ("thread-safe")
			wg1.Done()
		}()
	}

	wg1.Wait()
	fmt.Println("Counter1", counter1)

	var counter2 uint64
	var wg2 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func(num *uint64) {
			// even this approach produces inconsistent results
			*num += 1
		}(&counter2)
	}

	fmt.Println("Counter2", counter2)

}
