package main

import (
	"bufio"
	"fmt"
	"net/http"
)

var print = fmt.Println

func main() {
	print("-- Http Client Example --")

	resp, httpError := http.Get("https://gobyexample.com")

	if httpError != nil {
		panic(fmt.Sprintf("Network error: %v\n", httpError))
	}

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	// "for as long as there is a line to grab AND i is less than 5", aka "get me the first 5 lines"
	for i := 0; scanner.Scan() && i < 5; i++ {
		print(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
