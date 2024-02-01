package main

import "fmt"

var Reservations []Reservation = []Reservation{
	{
		Address:   "Belgrade",
		StartDate: "610/2023",
		EndDate:   "9/10/2023",
	},
	{
		Address:   "Novi Sad",
		StartDate: "3/10/2023",
		EndDate:   "19/10/2023",
	},
	{
		Address:   "Nis",
		StartDate: "5/10/2023",
		EndDate:   "6/10/2023",
	},
}

func CreateReservation(reservation Reservation) Reservation {
	fmt.Println("Created reservation:", reservation)

	return reservation
}

func UpdateReservation(reservation Reservation) Reservation {
	fmt.Println("Updated reservation:", reservation)
	return reservation
}
