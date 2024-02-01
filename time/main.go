package main

import (
	"fmt"
	"time"
)

var print = fmt.Println

func main() {
	print("-- Go Time --")
	now := time.Now()
	print(now)
	print("hello")
}
