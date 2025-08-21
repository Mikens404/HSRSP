package application

import (
	"context"

	"github.com/Mikens404/HSRSP/internal/domain"
)

func GetReservationSeat(ctx context.Context, trainNumber int, carNumber int) (domain.SeatReservationStatus, error) {
	return domain.SeatReservationStatus{}, nil
}
