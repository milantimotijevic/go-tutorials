package main

import (
	"fmt"
	"os"
	"text/template"
)

var print = fmt.Println

func main() {
	print("-- Go Text Templates --")
	template1 := template.New("template1")

	template1, err := template1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	template1.Execute(os.Stdout, "your mom")

	// shorthand for the above example error handling example
	template2 := template.New("template2")
	template.Must(template2.Parse("Value is {{.}}\n"))
	template2.Execute(os.Stdout, []int{1, 2, 3})
	template2.Execute(os.Stdout, "Moo")
}
