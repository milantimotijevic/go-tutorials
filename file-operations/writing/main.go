package main

import (
	"bufio"
	"fmt"
	"os"
)

var print = fmt.Println

func main() {
	print("-- File Writing --")

	// just dump stuff into a file
	writeErr := os.WriteFile("/tmp/sampletextfile.txt", []byte("dumped data"), 0644) // 0644 permissions
	print(writeErr)

	// if it already exists, it will be truncated
	file, _ := os.Create("/tmp/testingwrites.txt")
	defer file.Close()

	bytesWrittenCount, _ := file.Write([]byte{115, 111, 109, 101, 10})
	print(bytesWrittenCount)

	bytesWrittenCount2, _ := file.WriteString("Hello there what's up?!")
	print(bytesWrittenCount2)

	file.Sync() // ensure os has completed write operations

	bufferedWriter := bufio.NewWriter(file)
	bytesWrittenCount3, _ := bufferedWriter.WriteString("buffered string lol lol")
	print(bytesWrittenCount3)
	bufferedWriter.Flush() // tell os to complete the writing process

	// another "normal write"
	file.WriteString(" more text!")
	file.Sync()

	bufferedWriter.WriteString(" even more text OOMMG")
	bufferedWriter.Flush()
}
