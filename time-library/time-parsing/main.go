package main

import (
	"fmt"
	"time"
)

var print = fmt.Println

func main() {
	print("-- Time Parsing --")
	now := time.Now()
	print(now)
}
