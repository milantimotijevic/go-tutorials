package main

import (
	"fmt"
)

// remove specific item
func remove(list []string, item string) []string {
	for idx, v := range list {
		if v == item {
			list = append(list[:idx], list[idx+1:]...)
		}
	}

	return list
}

func main() {
	fmt.Println("Hello world!")
	var names []string = []string{"Mitlandir", "Andry", "Astarion", "Lae'zel"}
	fmt.Println(names)

	// careful, this will modify the underlying array, which can mess up your slice
	// remove(names, "Mitlandir")
	// that's why it's generally a good idea to reassign returned slices
	names = remove(names, "Mitlandir")
	fmt.Println(names)

	names = remove(names, "Shadowheart")
	// no change, never entered the loop
	fmt.Println(names)

	var numbers [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(numbers)
	// append(numbers, 5) // impossible, must be a slice
	numbers[2] = 555
	fmt.Println(numbers)
	// numbers[5] = 12212 // impossible, out of bounds

	fmt.Println("Two Dimensional Array:")
	var twoDee [2][2]int = [2][2]int{}

	for i := 0; i < len(twoDee); i++ {
		for j := 0; j < len(twoDee[i]); j++ {
			twoDee[i][j] = i + j
		}
	}

	fmt.Println(twoDee)

	fmt.Println("Two Dimensional Array Refactor Loop")

	twoDee = [2][2]int{}

	for _, item := range twoDee {
		fmt.Println(item)
	}

	var slicySlice []int = []int{1, 2, 3}
	fmt.Println(slicySlice)

	var slicySliceAllocated []int = make([]int, 3)
	fmt.Println(slicySliceAllocated)

	fmt.Println("--------")
	numSlice := []int{1, 2, 3, 4, 5, 6}
	// inclusive:exclusive
	anotherNumSlice := numSlice[2:4]
	fmt.Println(anotherNumSlice)
}
