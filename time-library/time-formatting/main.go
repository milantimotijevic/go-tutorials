package main

import (
	"fmt"
	"time"
)

var print = fmt.Println

func main() {
	print("-- Time Formating --")
	now := time.Now()
	print(now.Format(time.RFC3339))
	print(now.Format("3:04PM"))

	myBirthDate, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	print(myBirthDate)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	print(e)
}
