package main

import (
	"fmt"
	"sync"
	"time"
)

// TODO replace "any" with a proper generic
func announce[T any](items []T, wg *sync.WaitGroup) {
	// this means wg.Done() will be called when "announce" exits regardless of the outcome
	defer wg.Done()
	for i := range items {
		time.Sleep(1 * time.Second)
		fmt.Println(items[i])
	}

	// this is also possible, but make sure the code gets reached (best to use defer instead)
	// wg.Done()
}

func main() {
	fmt.Println("-- Go Wait Groups --")
	var numbers []int = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var names []string = []string{"Milan", "Milos", "Marko", "Jeca", "Neca", "Veca", "Mika"}

	// wait group is needed since I'm not using channels
	var wg sync.WaitGroup
	wg.Add(2)
	go announce(numbers, &wg)
	go announce(names, &wg)
	wg.Wait()

}
