package main

import (
	"fmt"
)

func initNextIntGetter() func() int {
	var counter int = 0
	return func() int {
		counter += 1
		return counter
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("-- Go Closures --")
	getNextInt := initNextIntGetter()
	fmt.Println(getNextInt())
	fmt.Println(getNextInt())
	fmt.Println(getNextInt())

	fmt.Println(factorial(5)) // 1 * 2 * 3 * 4 * 5

	var fibonacci func(num int) int
	// a recursive closure function has to be declared first like this
	fibonacci = func(num int) int {
		if num < 2 {
			return num
		}
		// it knows what to call because we declared it already
		return fibonacci(num-1) + fibonacci(num-2)
	}

	fmt.Println(fibonacci(7))
}
