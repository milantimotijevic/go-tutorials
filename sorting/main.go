package main

import (
	"fmt"
	"sort"
)

type StringSorter []string

// just a way to determine the length of the slice being sorted
func (slice StringSorter) Len() int {
	return len(slice)
}

func (slice StringSorter) Swap(first, second int) {
	slice[first], slice[second] = slice[second], slice[first]
}

func (slice StringSorter) Less(first, second int) bool {
	return len(slice[first]) < len(slice[second])
}

type Person struct {
	name string
	age  int
}

// byAge becomes a slice that can only take "instances" of Person
type byAgeSorter []Person

// just a way to determine the length of the slice being sorted (in this case a slice of persons)
func (slice byAgeSorter) Len() int {
	return len(slice)
}

func (slice byAgeSorter) Swap(first, second int) {
	slice[first], slice[second] = slice[second], slice[first]
}

func (slice byAgeSorter) Less(first, second int) bool {
	return slice[first].age < slice[second].age
}

func main() {
	fmt.Println("-- Sorting --")
	fmt.Println("-- Primitives --")
	numbers := []int{5, 10, 3, 1, 14}
	strings := []string{"c", "b", "a", "cat", "moose", "loose", "caboose"}

	sort.Ints(numbers)
	sort.Strings(strings)

	fmt.Println(numbers)
	fmt.Println(strings)
	fmt.Println("---------")

	fmt.Println("-- Sorting By Sorter Functions --")
	names := []string{"Pera", "Perica", "Petronije", "Mika", "Mikla", "Milka"}
	sort.Sort(StringSorter(names))
	fmt.Println(names)
	fmt.Println("-------------")

	fmt.Println("-- Sorting Structs --")
	people := []Person{
		{"Pera", 15}, {"Mika", 12}, {"Zika", 20},
	}

	sort.Sort(byAgeSorter(people))
	fmt.Println(people)
}
