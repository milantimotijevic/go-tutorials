package main

import (
	"fmt"
	"strconv"
)

var print = fmt.Println

func main() {
	print("-- Parsing Numbers --")
	print(strconv.ParseInt("5", 0, 64))
	print(strconv.ParseFloat("5.32", 64))
}
