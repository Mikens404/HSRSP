package application

import (
	"context"

	"github.com/Mikens404/HSRSP/internal/domain"
)

type ReservationService interface {
	GetReservationSeat(ctx context.Context, trainNumber int, carNumber int) (domain.SeatReservationStatus, error)
}

type reservationService struct {
	trainRepository       domain.TrainRepository
	reservationRepository domain.ReservationRepository
}

func NewReservationService(
	trainRepository domain.TrainRepository,
	reservationRepository domain.ReservationRepository,
) ReservationService {
	return &reservationService{
		trainRepository:       trainRepository,
		reservationRepository: reservationRepository,
	}
}

func (r *reservationService) CreateReservation() {

}

func (r *reservationService) GetReservationSeat(ctx context.Context, trainNumber int, carNumber int) (domain.SeatReservationStatus, error) {
	return domain.SeatReservationStatus{}, nil
}
