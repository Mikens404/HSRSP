package application

import (
	"context"
	"errors"

	"github.com/Mikens404/HSRSP/internal/domain"
)

/* type ReservationService interface {
	CreateReservation(ctx context.Context, createReservationParams domain.CreateReservationParams) error
	GetReservationSeat(ctx context.Context, trainNumber int, carNumber int) (domain.SeatReservationStatus, error)
} */

type ReservationService struct {
	reservationRepository domain.ReservationRepository
}

func NewReservationService(
	reservationRepository domain.ReservationRepository,
) ReservationService {
	return ReservationService{
		reservationRepository: reservationRepository,
	}
}

func (r ReservationService) CreateReservation(ctx context.Context, createReservationParams domain.CreateReservationParams) error {
	if createReservationParams.ReservationPeople != len(createReservationParams.ReservationSeats) {
		return errors.New("予約人数と予約席数が一致しません") // TODO:エラーを作る
	}
	if err := r.reservationRepository.InsertReservation(ctx, createReservationParams); err != nil {
		return err
	}
	return nil
}

func (r ReservationService) GetReservationSeat(ctx context.Context, trainNumber int, carNumber int) (domain.SeatReservationStatus, error) {
	return domain.SeatReservationStatus{}, nil
}
