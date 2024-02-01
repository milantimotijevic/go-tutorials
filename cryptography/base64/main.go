package main

import (
	"encoding/base64"
	"fmt"
)

var print = fmt.Println

func main() {
	print("-- Base64 --")
	someString := "LaGenteEstaMuyLoca"

	encode := base64.StdEncoding.EncodeToString([]byte(someString))
	print(encode)

	decode, _ := base64.StdEncoding.DecodeString(encode)
	print(string(decode))

	// url-friendly encoding approach, I guess...
	urlEncode := base64.URLEncoding.EncodeToString([]byte(someString))
	print(urlEncode)

	urlDecode, _ := base64.URLEncoding.DecodeString(urlEncode)
	print(string(urlDecode))
}
