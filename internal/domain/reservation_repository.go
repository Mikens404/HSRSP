package domain

import (
	"context"
)

type ReservationRepository interface {
	CreateReservation(ctx context.Context, createReservationParams CreateReservationParams) error
	FindReservation(ctx context.Context, reservationNumber int) (Reservation, error)
	UpdateReservation(ctx context.Context, updateReservationParams UpdateReservationParams) error
	FindReservationSeats(ctx context.Context, trainNumber int, carNumber int) (SeatReservationStatus, error)
}
