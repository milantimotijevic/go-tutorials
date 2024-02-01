package main

import (
	"fmt"
	"os"
	"strings"
)

var print = fmt.Println

func main() {
	print("-- Environment Variables Examples --")
	print(os.Getenv("cat"))
	os.Setenv("dog", "woof")
	print(os.Getenv("dog"))

	for _, envVar := range os.Environ() {
		pair := strings.SplitN(envVar, "=", 2) // 2 is the max num of elements to return
		print(pair)
	}
}
