package main

import (
	"fmt"
	"math/rand"
	"time"
)

var print = fmt.Println

func main() {
	print("-- RNG --")
	rand.Seed(time.Now().UnixNano()) // seeds globally
	print(rand.Intn(3) + 1)          // 1-3

	source1 := rand.NewSource(time.Now().UnixMicro())
	rng1 := rand.New(source1) // new generator with its own seed (source)
	print(rng1.Intn(3) + 1)

	rng2 := rand.New(rand.NewSource(567)) // hardcoded source will always produce the same num
	// if the calling arg is the same
	print(rng2.Intn(222))
}
