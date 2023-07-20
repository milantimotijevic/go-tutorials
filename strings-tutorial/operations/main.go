package main

import (
	"fmt"
	s "strings" // alias it for ease of use
)

var p = fmt.Println // alias this one as well

func main() {
	p(s.Contains("cat", "c"))         // true
	p(s.Count("catt", "t"))           // 2
	p(s.HasPrefix("cattery", "cat"))  // true
	p(s.HasSuffix("cattery", "tery")) // true
	p(s.Index("catccc", "c"))         // 0 // index of first encountered
	p(s.Join([]string{
		"what",
		"does",
		"the",
		"fox",
		"say",
		"?"},
		"-")) // it's obv....
	p(s.Replace("moo", "o", "L", 1))      // mLo // last arg means how many to replace
	p(s.Replace("moooooo", "o", "L", -1)) // negative last arg means replace all
	p(s.Split("cat goes meow", "goes"))   // ["cat", "meow"]
	p(s.ToLower("OMG YOU NOOB"))
	p(s.ToUpper("many whelps, handle it!"))

	print("SplitN example:")
	sampleString := "cat=meow=dog=woof"
	// extract a maximum of 3 elements (place remainder in the last one)
	splitResult := s.SplitN(sampleString, "=", 3)
	print("\n")
	for _, item := range splitResult {
		print(item, "\n")
	}
}
