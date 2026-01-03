package application

import (
	"context"
	"errors"

	"github.com/Mikens404/HSRSP/internal/domain"
)

type ReservationService interface {
	CreateReservation(ctx context.Context, createReservationParams domain.CreateReservationParams) error
	GetReservationSeat(ctx context.Context, trainNumber int, carNumber int) (domain.SeatReservationStatus, error)
}

type reservationServiceImpl struct {
	reservationRepository domain.ReservationRepository
}

func NewReservationService(
	reservationRepository domain.ReservationRepository,
) ReservationService {
	return &reservationServiceImpl{
		reservationRepository: reservationRepository,
	}
}

func (r *reservationServiceImpl) CreateReservation(ctx context.Context, createReservationParams domain.CreateReservationParams) error {
	if createReservationParams.ReservationPeople != len(createReservationParams.ReservationSeats) {
		return errors.Join(domain.ErrNotmatchReservationNumber)
	}
	if err := r.reservationRepository.InsertReservation(ctx, createReservationParams); err != nil {
		return err
	}
	return nil
}

func (r *reservationServiceImpl) GetReservationSeat(ctx context.Context, trainNumber int, carNumber int) (domain.SeatReservationStatus, error) {
	return domain.SeatReservationStatus{}, nil
}
