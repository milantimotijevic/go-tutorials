package main

import (
	"fmt"
	"package-structure/services"
)

func main() {
	fmt.Println("Hello")

	// it's a bit clunky if we want to put all services in the same package
	// we have to instantiate them before calling their methods
	// probably better to just have a separate folder/package for each service
	reservationService := services.ReservationService{}
	str := reservationService.CreateReservation()

	fmt.Println(str)
}
