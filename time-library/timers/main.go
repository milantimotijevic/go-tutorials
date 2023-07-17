package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-- Go Time --")

	timer1 := time.NewTimer(time.Second * 2)

	fmt.Println(<-timer1.C) //.C stands for channel, which returns the time it was fired

	timer1Stopped := timer1.Stop()
	// it wasn't stopped, so this will be false
	fmt.Println("Timer 1 was stopped:", timer1Stopped)

	timer2 := time.NewTimer(time.Second * 2)
	timer2Stopped := timer2.Stop()
	// fmt.Println(<-timer2.C) // this would panic, because nothing would ever be sent to this channel
	fmt.Println("Timer 2 was stopped:", timer2Stopped)

	fmt.Println("Time to sleep")
	time.Sleep(time.Second)
	fmt.Println("Sleep is over")

	fmt.Println("time.After example")
	// returns a channel that will return a time object after a set amount of time
	afterTimer := time.After(time.Second * 3)
	fmt.Println("time.After result:", <-afterTimer)

	fmt.Println("Yet another time")
	yetAnotherTimer := time.NewTimer(time.Second * 3)
	fmt.Println(yetAnotherTimer)
}
