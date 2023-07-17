package main

import (
	"fmt"
)

func getSum(nums ...int) int {
	var total = 0
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	fmt.Println("-- Go Variadic Functions --")
	fmt.Println(getSum(2))
	fmt.Println(getSum(2, 2))
	fmt.Println(getSum(5, 5))
	var sliceOfNums []int = []int{1, 2, 3}
	fmt.Println(getSum(sliceOfNums...))
}
