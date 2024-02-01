package main

import (
	"fmt"
	"time"
)

var print = fmt.Println

func main() {
	print("-- Go Time --")
	now := time.Now()
	print(now)

	myBirthDate := time.Date(1989, 12, 10, 18, 30, 25, 0, time.UTC)
	print(myBirthDate)

	print(myBirthDate.Year())
	print(myBirthDate.Month())
	print(myBirthDate.Day())
	print(myBirthDate.Hour())
	print(myBirthDate.Minute())
	print(myBirthDate.Second())
	print(myBirthDate.Nanosecond())
	print(myBirthDate.Location())

	print(myBirthDate.Weekday())

	print(myBirthDate.Before(now))
	print(myBirthDate.After(now))
	print(myBirthDate.Equal(now))

	diffBdayNow := myBirthDate.Sub(now)
	print(diffBdayNow)

	diffNowBday := now.Sub(myBirthDate)
	print(diffNowBday)

	print(myBirthDate.Add(diffNowBday))
	print(myBirthDate.Add(-diffNowBday))

	print(now.Unix())
	print(now.UnixMilli())
	print(now.UnixNano())
	print(time.Unix(now.Unix(), 0))
	print(time.Unix(0, now.UnixNano()))
}
