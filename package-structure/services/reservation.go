package services

type ReservationService struct{}

func (r *ReservationService) GetReservation() string {
	return "got reservation"
}

func (r *ReservationService) CreateReservation() string {
	return "created reservation"
}
