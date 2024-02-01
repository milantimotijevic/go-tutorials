package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mutex    sync.Mutex
	counters map[string]int
}

func (c *Container) incrementCounter(counterName string) {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	c.counters[counterName]++
}

func main() {
	fmt.Println("-- Mutexes --")
	container := Container{
		counters: map[string]int{
			"first":  0,
			"second": 0,
		},
	}

	var wg sync.WaitGroup
	doIncrement := func(counterName string) {
		// since we have this many iterations, a syncing mechanism is necessary
		for i := 0; i < 10000; i++ {
			container.incrementCounter(counterName)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("first")
	go doIncrement("first")
	go doIncrement("second")

	wg.Wait()

	fmt.Println("First:", container.counters["first"])
	fmt.Println("Second:", container.counters["second"])
}
