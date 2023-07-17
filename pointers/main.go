package main

import "fmt"

func simpleByValue(age int, name string) {
	fmt.Println(age)
	fmt.Println(name)

	// change the value, the change is only local
	age = 34
	name = "Milan Timotijevic"

	// the change is reflected in this scope
	fmt.Println(age)
	fmt.Println(name)
}

func simpleByReference(agePtr *int, namePtr *string) {
	// print the pointers themselves
	fmt.Println(agePtr)
	fmt.Println(namePtr)

	// print it dereferenced (so we get the actual value getting pointed at)
	fmt.Println(*agePtr)
	fmt.Println(*namePtr)

	// dereference and change the values
	*agePtr = 34
	*namePtr = "Milan Timotijevic"
	/*
	*namePtr = "super " + *namePtr // to use contents of the original string
	 */

	// the values are changed here, but also in all other scopes
	fmt.Println(*agePtr)
	fmt.Println(*namePtr)
}

type Person struct {
	name string
	age  int
}

// pass object by value
func objectByValue(person Person) {
	fmt.Println(person)
	// change only in this context
	person.name = "Milan Timotijevic"
	person.age = 34
	fmt.Println(person)
}

// pass object by reference
func objectByReference(person *Person) {
	// gets automatically dereferenced
	fmt.Println(person)
	// can dereference it to get output with "&"
	fmt.Println(*person)

	// can access its fields without the need to dereference (it's not possible, in fact)
	fmt.Println(person.name)

	person.name = "Milan Timotijevic"
	person.age = 34

	fmt.Println(person)
}

// objects and arrays follow the same rules
func arrayByValue(people [2]string) {
	fmt.Println(people)
	people[0] = "Milan Timotijevic"
	fmt.Println(people)
}

func arrayByReference(peoplePtr *[2]string) {
	fmt.Println(peoplePtr)
	fmt.Println(*peoplePtr)
	// it automatically gets dereferenced, same as object
	peoplePtr[0] = "Milan Timotijevic"
	fmt.Println(*peoplePtr)
}

// slices and maps follow different rules from objects and arrays
func sliceByValue(names []string) {
	// looks like slices get passed by reference by default
	fmt.Println(names)
	names[0] = "Milan Timotijevic"
	fmt.Println(names)
}

func mapByValue(person map[string]string) {
	fmt.Println(person)
	person["firstName"] = "MILAN"
	fmt.Println(person)
}

func mapCompletelyChangeByValue(person map[string]string) {
	fmt.Println(person)
	person = map[string]string{"status": "changed"}
	fmt.Println(person)
}

func mapCompletelyChangeByReference(person *map[string]string) {
	fmt.Println(person)
	*person = map[string]string{"status": "changed"}
	fmt.Println(person)
}

/**
Takeaway: slices and maps are passed by reference by default, so modifying them in the receiving function
will modify them everywhere else. You can still pass their reference directly, which will allow you
to replace the entire map/slice

All other data types are passed by value, which means they don't get modified outside of the
receiving function, unless if passed by reference explicitly

If you want to change the whole thing entirely (i.e. replace the object/map/slice/etc.) in the receiving
function, pass it as a reference
*/
func main() {
	/*
		var age int = 33
		var name string = "Milan"
		var agePtr *int = &age
		var namePtr *string = &name

		fmt.Println(agePtr)
		fmt.Println(namePtr)
		fmt.Println(name)

		// create a reference to a string value
		var fullNamePtr *string = namePtr
		fmt.Println(fullNamePtr)
		fmt.Println(*fullNamePtr)
		// dereference the pointer back onto the value it points to
		*fullNamePtr = "Milan Timotijevic"
		fmt.Println(*fullNamePtr)

		*agePtr = 34
		fmt.Println(*agePtr)
		// the values are changed regardless of which variable we access it from
		fmt.Println(age)
		fmt.Println(name)
	*/

	/*
		var age int = 33
		var name string = "Milan"
		fmt.Println(age)
		fmt.Println(name)
		// pass by value and have the value changed in that function
		simpleByValue(age, name)
		// the value is unchanged in this scope
		fmt.Println(age)
		fmt.Println(name)
	*/
	/*
		var age int = 33
		var name string = "Milan"
		fmt.Println(age)
		fmt.Println(name)
		// pass by reference and have the value the reference points to changed
		simpleByReference(&age, &name)
		// the value is changed even in this scope
		fmt.Println(age)
		fmt.Println(name)
	*/

	/*
		var person Person = Person{
			name: "Milan",
			age:  33,
		}

		fmt.Println(person)
		objectByValue(person)
		// unchanged in this context, even though it got changed in helper's context
		fmt.Println(person)
	*/
	/*
		var person Person = Person{
			name: "Milan",
			age:  33,
		}

		fmt.Println(person)
		objectByReference(&person)
		// changed even in this context
		fmt.Println(person)
	*/
	/*
		var people [2]string = [2]string{"Milan", "Milos"}
		fmt.Println(people)
		arrayByValue(people)
		// value not changed in this context
		fmt.Println(people)
	*/
	/*
		var people [2]string = [2]string{"Milan", "Milos"}
		fmt.Println(people)
		arrayByReference(&people)
		// gets changed even in this context
		fmt.Println(people)
	*/
	/*
		var names = []string{"Milan", "Milos"}
		fmt.Println(names)
		sliceByValue(names)
		// it gets changed even in this context
		fmt.Println(names)
	*/

	/*
		var person map[string]string = map[string]string{}
		person["firstName"] = "Milan"
		fmt.Println(person)
		mapByValue(person)
		// it gets changed even in this context
		fmt.Println(person)
	*/

	/*
		var person map[string]string = map[string]string{"firstName": "Milan"}
		fmt.Println(person)
		mapCompletelyChangeByValue(person)
		// not changed in this context
		fmt.Println(person)
	*/

	/*
		var person map[string]string = map[string]string{"firstName": "Milan"}
		fmt.Println(person)
		mapCompletelyChangeByReference(&person)
		// changed in this context
		fmt.Println(person)
	*/
}
