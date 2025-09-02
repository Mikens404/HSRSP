package application

import (
	"context"
	"errors"

	"github.com/Mikens404/HSRSP/internal/domain"
)

type ReservationService interface {
	GetReservationSeat(ctx context.Context, trainNumber int, carNumber int) (domain.SeatReservationStatus, error)
}

type reservationService struct {
	reservationRepository domain.ReservationRepository
}

func NewReservationService(

	reservationRepository domain.ReservationRepository,
) ReservationService {
	return &reservationService{

		reservationRepository: reservationRepository,
	}
}

func (r *reservationService) CreateReservation(ctx context.Context, createReservationParams domain.CreateReservationParams) error {
	if createReservationParams.ReservationPeople != len(createReservationParams.ReservationSeats) {
		return errors.New("")
	}
	if err := r.reservationRepository.InsertReservation(ctx, createReservationParams); err != nil {
		return err
	}
	return nil
}

func (r *reservationService) GetReservationSeat(ctx context.Context, trainNumber int, carNumber int) (domain.SeatReservationStatus, error) {
	return domain.SeatReservationStatus{}, nil
}
