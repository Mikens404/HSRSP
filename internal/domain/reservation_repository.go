package domain

import (
	"context"
)

type ReservationRepository interface {
	CreateReservation(
		ctx context.Context,
		reservationTrainNumber int,
		boardingStation string,
		getOffStation string,
		reservationSeatList []ReservationSeat,
		reservationPeople int,
		customerInfo string,
	) error
	FindReservation(ctx context.Context, reservationNumber int) (Reservation, error)
	UpdateReservation(
		ctx context.Context,
		reservationTrainNumber int,
		boardingStation string,
		getOffStation string,
		reservationSeatList []ReservationSeat,
		reservationPeople int,
		reservationStatus CustomerReservationStatus,
		customerInfo string,
	) error
	FindReservationSeats(ctx context.Context, trainNumber int, carNumber int) (SeatReservationStatus, error)
}
