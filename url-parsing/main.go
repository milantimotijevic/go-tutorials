package main

import (
	"fmt"
	"net"
	"net/url"
)

var print = fmt.Println

func main() {
	print("-- URL Parsing --")
	urlString1 := "postgres://user:pass@host.com:5432/path?k=v#f"

	url1, _ := url.Parse(urlString1)

	print(url1.Scheme)
	print(url1.User)
	print(url1.User.Username())
	pass1, _ := url1.User.Password()
	print(pass1)
	print(url1.Host)
	host1, port1, _ := net.SplitHostPort(url1.Host)
	print("host is", host1, "port is", port1)

	print(url1.Path)
	print(url1.Fragment) // after the hash
	print(url1.RawQuery)

	query1, _ := url.ParseQuery(url1.RawQuery) // returns a map
	print(query1)
}
