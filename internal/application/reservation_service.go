package application

import (
	"context"

	"github.com/Mikens404/HSRSP/internal/domain"
)

type ReservationService interface {
	GetReservationSeat(ctx context.Context, trainNumber int, carNumber int) (domain.SeatReservationStatus, error)
}

/*
	 type reservationServiceImpl struct {
		reservationRepository domain.ReservationRepository
	}
*/
func GetReservationSeat(ctx context.Context, trainNumber int, carNumber int) (domain.SeatReservationStatus, error) {
	return domain.SeatReservationStatus{}, nil
}
