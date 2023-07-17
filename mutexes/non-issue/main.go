package main

import (
	"fmt"
	"sync"
)

type Container struct {
	counters map[string]int
}

func (c *Container) incrementCounter(counterName string) {
	c.counters[counterName]++
}

func main() {
	container := Container{
		counters: map[string]int{
			"first":  0,
			"second": 0,
		},
	}

	var wg sync.WaitGroup
	doIncrement := func(counterName string) {
		/*
		We probably won't encounter a problem here because there are so few iterations that the
		likelihood of a concurrent map write is close to non-existent
		However, if we upped the number to 10000, we would get "fatal error: concurrent map writes"
		*/
		for i := 0; i < 1000; i++ {
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
