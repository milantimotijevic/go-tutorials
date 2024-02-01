package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	p := filepath.Join("dir1", "dir2", "dir3", "filename.txt")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p)) // preserves the suffix
	// trim the suffix; args are filename.txt and txt
	fmt.Println("Trimmed suffix", strings.TrimSuffix(filepath.Base(p), filepath.Ext(filepath.Base(p))))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext))

	// how to get from first to the other
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	fmt.Println(os.Args[0]) // name of the binary being executed
}
