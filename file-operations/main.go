package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var print = fmt.Println

func main() {
	print("-- Reading Files --")

	print("-- Read Entire File --")
	readResult1, _ := os.ReadFile("/tmp/shopping-list.txt") // read everything
	print(string(readResult1))

	// read successively
	print("-- Partially Read --")
	file1, _ := os.Open("/tmp/shopping-list.txt")
	byteSlice1 := make([]byte, 10)                 // read up to 10 bytes
	actuallyReadCount, _ := file1.Read(byteSlice1) // byteSlice1 now holds those characters
	print(string(byteSlice1[:actuallyReadCount]))
	// print(string(byteSlice1)) // or simply like this...

	print("-- Seeking --")
	// "move the pointer (offset) two characters forward (positive number), from
	// 0 - the start of the file, 1 - the current offset, 2 - the end of the file"
	offset1, _ := file1.Seek(2, 0)
	print("offset1 is", offset1)
	// this will be applied to every subsequent read
	byteSlice2 := make([]byte, 10) // slice size will also limit it
	file1.Read(byteSlice2)
	print(string(byteSlice2))

	print("-- Seeking And Reading At Least n --")
	file1.Seek(6, 0)
	byteSlice3 := make([]byte, 5)
	// will read at least 2 bytes and throw an error if there are fewer
	actuallyReadCount2, _ := io.ReadAtLeast(file1, byteSlice3, 2)
	print("actuallyReadCount2", actuallyReadCount2)
	// still respects slice size and offset
	print(string(byteSlice3))
	file1.Seek(0, 0) // rewind to start

	print("-- Buffered Reader --")
	bufferedReader1 := bufio.NewReader(file1)
	buffReadResult1, _ := bufferedReader1.Peek(2) // read ahead but without needing to move the offset
	print(string(buffReadResult1))
	// it also buffers data internally so it improves performance, I guess.......

	file1.Close()
}
