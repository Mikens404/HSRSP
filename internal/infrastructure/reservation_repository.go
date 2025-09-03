package infrastructure

import (
	"context"

	"github.com/Mikens404/HSRSP/internal/domain"
)

// いつかDB設計する😭

type reservationRepositoryImpl struct{}

func NewReservationRepository() trainRepositoryImpl {
	return trainRepositoryImpl{}
}

func (r reservationRepositoryImpl) InsertReservation(ctx context.Context, createReservationParams domain.CreateReservationParams) error {
	return nil
}

func (r reservationRepositoryImpl) FindReservation(ctx context.Context, reservationNumber int) (domain.Reservation, error) {
	return domain.Reservation{}, nil
}

func (r reservationRepositoryImpl) UpdateReservation(ctx context.Context, updateReservationParams domain.UpdateReservationParams) error {
	return nil
}

func (r reservationRepositoryImpl) FindReservationSeats(ctx context.Context, trainNumber int, carNumber int) (domain.SeatReservationStatus, error) {
	return domain.SeatReservationStatus{}, nil
}
