package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("-- Atomic Counter Mutex Example --")

	var counter int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer mutex.Unlock()
			// has to be the same mutex object across all goroutines that we want to sync
			// once it encounters Lock(), it checks whether it's already locked
			// if it is, the routine is removed from the execution pool (until the mutex is unlocked)
			// if it isn't, the routine places a lock on it and proceeds with the code
			mutex.Lock()
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter", counter)
}
