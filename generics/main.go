package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// go mod tidy was needed to ensure above dependency was available
func powerUp[T constraints.Integer | constraints.Float](first T, second T) T {
	return (first + second) * 2
}

func checkType(item interface{}) {
	// .(type) is only usable in a switch statement because Go forces you to follow good coding practices
	// v gets the value of the argument, whereas case values cannot be directly assigned to a variable
	switch v := item.(type) {
	case int:
		fmt.Println("The parameter is an integer:", v)
	case string:
		fmt.Println("The parameter is a string:", v)
	default:
		fmt.Println("The parameter has an unknown type:", v)
	}
}

func checkTypeGeneric[T any](item T) {
	// have to cast "item" to "any" because it's a generic type
	switch any(item).(type) {
	case int:
		fmt.Println("It's a freaking int")
	case float32:
		fmt.Println("It's a freaking float32")
	case string:
		fmt.Println("It's a freaking string")
	default:
		fmt.Println("I don't freaking know!")
	}
}

type Character struct {
	race  string
	name  string
	class string
	age   int
}

type Ability struct {
	name     string
	cost     int
	strength int
}

type GameItem interface {
	GetName() string
}

func (c Character) GetName() string {
	return c.name
}

func (a Ability) GetName() string {
	return a.name
}

func announceName(item GameItem) {
	fmt.Println(item.GetName())
}

// this wouldn't actually work because you can't compare different types of interfaces
// same would probably apply to using "any" (it's supposedly also an interface type)
func includesInterface(list []interface{}, item any) bool {
	for _, v := range list {
		if item == v {
			return true
		}
	}

	return false
}

// generically check whether passed list contains passed item
func includes[T Ability | Character | string](list []T, item T) (bool, int) {
	for idx, v := range list {
		if item == v {
			return true, idx
		}
	}

	return false, -1
}

func push[T Ability | Character | string](list []T, item T) []T {
	list = append(list, item)
	return list
}

// directly modify the slice through its pointer
// probably wouldn't do this, it's not very readable
func pushByPtr[T Ability | Character | string](list *[]T, item T) {
	*list = append(*list, item)
}

// remove specific item
func remove[T Ability | Character | string](list []T, item T) []T {
	for idx, v := range list {
		if v == item {
			list = append(list[:idx], list[idx+1:]...)
		}
	}

	return list
}

func removeLast[T Ability | Character | string](list []T) ([]T, T) {
	removed := list[len(list)-1]
	list = append(list[:len(list)-1])

	return list, removed
}

func removeFirst[T Ability | Character | string](list []T) ([]T, T) {
	removed := list[0]
	list = append(list[1:])

	return list, removed
}

/*
// something like this... I think... (doesn't work right now)
func removeByPtr[T Ability | Character](list *[]T, item T) {
	for idx, v := range *list {
		if v == item {
			*list = append(list[:idx], list[idx+1:]...)
		}
	}
}
*/

func main() {
	player := Character{
		race:  "High Elf",
		name:  "Mitlandir",
		class: "Wizard",
		age:   33,
	}

	hobgoblin := Character{
		race:  "Hobgoblin",
		name:  "Rgzuzgeenzg",
		class: "Fighter",
		age:   20,
	}

	laeZel := Character{
		race:  "Githyanki",
		name:  "Lae'zel",
		class: "Fighter",
		age:   35,
	}

	astarion := Character{
		race:  "High Elf",
		name:  "Astarion",
		class: "Rogue",
		age:   352,
	}

	shadowheart := Character{
		race:  "High Half-Elf",
		name:  "Shadowheart",
		class: "Cleric",
		age:   32,
	}

	fireBolt := Ability{
		name:     "Fire Bolt",
		cost:     1,
		strength: 8,
	}

	rayOfFrost := Ability{
		name:     "Ray of Frost",
		cost:     1,
		strength: 6,
	}

	hideousLaughter := Ability{
		name:     "Hideous Laughter",
		cost:     2,
		strength: 0,
	}
	/*
		grease := Ability{
			name:     "Grease",
			cost:     2,
			strength: 0,
		}

		moonBeam := Ability{
			name:     "Moon Beam",
			cost:     2,
			strength: 13,
		}
	*/
	var characters []Character = []Character{player, hobgoblin}
	var abilities []Ability = []Ability{fireBolt, rayOfFrost}

	fmt.Println(includes(characters, player))
	fmt.Println(includes(abilities, rayOfFrost))
	fmt.Println(includes(abilities, hideousLaughter))
	// does not include the last one

	abilities = push(abilities, hideousLaughter)
	fmt.Println(includes(abilities, hideousLaughter))
	// now it includes it

	// pass the slice by pointer, which means we don't need to return and recapture the slice
	pushByPtr(&abilities, Ability{
		name:     "Sleep",
		cost:     2,
		strength: 0,
	})
	fmt.Println(abilities)

	// capture the returned value (I don't know if I could write it with a pointer...)
	abilities = remove(abilities, hideousLaughter)
	fmt.Println(abilities)

	fmt.Println("*********")
	fmt.Println("String time!")
	var strings []string = []string{"pera"}
	fmt.Println(strings)
	fmt.Println(includes(strings, "pera"))
	fmt.Println(includes(strings, "mika"))
	strings = push(strings, "mika")
	fmt.Println(strings)
	strings = remove(strings, "pera")
	// calling "remove" without reassigning returned value messes up the slice
	// this is because the underlying array gets modified
	fmt.Println(strings)

	fmt.Println("****************")
	fmt.Println("Queue time!")
	var characterQueue []Character = []Character{}
	fmt.Println(characterQueue)
	characterQueue = push(characterQueue, player)
	fmt.Println(characterQueue)
	characterQueue = push(characterQueue, laeZel)
	characterQueue = push(characterQueue, astarion)
	characterQueue = push(characterQueue, shadowheart)
	fmt.Println(characterQueue)
	characterQueue, lastRemoved := removeLast(characterQueue)
	fmt.Printf("Queue: %v, last removed removed item: %v\n", characterQueue, lastRemoved)

	characterQueue, firstRemoved := removeFirst(characterQueue)
	fmt.Printf("Queue: %v, first removed removed item: %v\n", characterQueue, firstRemoved)
}

/*
Can we create a slice of generic types?
Can we pass a generic type to a function and then "unpack" it as a non-generic and call its methods?
*/
