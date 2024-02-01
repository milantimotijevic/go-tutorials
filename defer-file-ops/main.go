package main

import (
	"fmt"
	"os"
)

func createFile(path string) *os.File {
	fmt.Println("creating")
	file, error := os.Create(path)
	if error != nil {
		panic(fmt.Sprintf("Error creating file at location %v\n", path))
	}

	return file
}

func writeFile(file *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(file, "hello")
}

func closeFile(file *os.File) {
	fmt.Println("closing")
	err := file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error %v\n", err)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("-- Defer and File Ops --")

	file := createFile("/tmp/defer-example.txt")
	defer closeFile(file)
	writeFile(file)
}
