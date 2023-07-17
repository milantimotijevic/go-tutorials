package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var print = fmt.Println

func main() {
	print("-- Go Regexp --")
	// simple matching can be done directly off of regexp package
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	print(match)

	// for most other methods it's needed to compile first
	compileResult, err := regexp.Compile("p([a-z]+)ch")
	if err != nil {
		panic("Yeah, no, that regexp ain't compilin...")
	}

	// shorthand for above err handling
	compileResult = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", compileResult)

	print(compileResult.MatchString("peach"))
	print(compileResult.FindString("peach punch"))
	print(compileResult.FindStringIndex("peach punch"))
	// match both the whole regexp and the submatch ([a-z]+) // [peach ea]
	print(compileResult.FindStringSubmatch("peach punch"))
	// return n strings that match the regexp; negative number for all
	print(compileResult.FindAllString("peach punch pinch", -1))
	// thus one returns a multi-dimensional slice; I don't know.. I just don't... Maybe char indices?
	print(compileResult.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// apparently, you can also use bytes instead of strings.. that's useful because... errr... it is?
	print(compileResult.Match([]byte("peach")))

	// replaces the part that matches the regexp?
	print(compileResult.ReplaceAllString("a peach", "<fruit>"))

	// can also manipulate bytes, because that's totally something I'd do...
	// this first section creates a slice of bytes that correspond to given string (in UTF-8), I think...
	in := []byte("a peach")
	out := compileResult.ReplaceAllFunc(in, bytes.ToUpper)
	print(out)
	print(string(out))
}
