package main

import (
	"fmt"
)

var print = fmt.Println

func main() {
	print("-- Math Helper Unit Testing --")
}

func Sum(a int, b int) int {
	return a + b
}

func Multiply(a int, b int) int {
	return a * b
}

func Subtract(a int, b int) int {
	return a - b
}

func Divide(a int, b int) int {
	return a / b
}
