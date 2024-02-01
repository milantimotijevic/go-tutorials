package main

type Reservation struct {
	Address   string `json:"address" validate:"required"`
	StartDate string `json:"startDate" validate:"required"`
	EndDate   string `json:"endDate" validate:"required"`
}
