package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func processObject(person Person) {
	// You sent me an object, I copied it and attached my own pointer to the copy 0:-)
	person.name = "Milan Timotijevic"
}

func processObjectByPtr(person *Person) {
	// You sent me an object pointer, I took that pointer and now it's mine! Screw you!
	person.name = "Milan Timotijevic"

	// If I wanted, I could make it point to something else and you could woudn't be able to do
	// a DAMN thing about it!
	/*
		*person = Person{
			name: "Pera Peric",
			age: 20,
		}
	*/
}

// this is a copy of a struct containing a ptr to underlying array, length and capacity
func processSlice(numbers []int) {
	// a slice is still an "object" (struct)
	// it goes directly to the underlying array and changes its nth element
	numbers[0] = 9999

	// I can't completely screw you up...
	numbers = []int{5, 6, 7}
	// this is now a new "object" with a new underlying array
}

// this is a pointer to the struct containing a ptr to underlying array, length and capacity
func processSliceByPtr(numbers *[]int) {
	// I can easily do this
	(*numbers)[0] = 7777777

	// But why stop there when I can COMPETELY screw you up??
	*numbers = []int{9999, 1231231}
}

// this is ALWAYS a pointer to a hash map
func processMap(m map[string]string) {
	m["Shadowheart"] = "Cleric"
}

// if I add * it becomes a pointer to a pointer to a hash map
func processMapByPointerToPointerToHashMap(m *map[string]string) {
	*m = map[string]string{
		"Cat": "Meow",
	}
}

func main() {
	fmt.Println("-- Go Pointers Refreshers --")

	person := Person{
		name: "Milan",
		age:  33,
	}

	processObject(person)
	fmt.Println("after processObject: ", person)

	processObjectByPtr(&person)
	fmt.Println("after processObjectByPtr: ", person)

	numbers := []int{1, 2, 3}
	// a slice contains a ptr to the underlying array, length and cap
	/* for example
	{
		length: 5,
		cap: 10,
		ptrToArray: 1xsldjsldfjsdl
	}
	*/

	processSlice(numbers)
	fmt.Println("after processSlice: ", numbers)

	processSliceByPtr(&numbers)
	fmt.Println("after processSliceByPtr: ", numbers)

	var companionsMap map[string]string = map[string]string{
		"Gale":    "Wizard",
		"Lae'Zel": "Fighter",
	}

	processMap(companionsMap)
	fmt.Println(companionsMap)

	processMapByPointerToPointerToHashMap(&companionsMap)
	fmt.Println(companionsMap)
}
