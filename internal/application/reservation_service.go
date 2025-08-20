package application

import (
	"context"

	"github.com/Mikens404/HSRSP/internal/domain"
)

func GetReservationSeat(ctx context.Context, trainNumber int, carNumber int) []domain.SeatReservationStatus {
	return []domain.SeatReservationStatus{}
}
