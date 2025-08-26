package application

import (
	"context"

	"github.com/Mikens404/HSRSP/internal/domain"
)

type TrainService interface {
	GetTrainInfo(ctx context.Context, trainNumber int) (domain.Train, error)
}

type trainService struct {
	trainRepository       domain.TrainRepository
	reservationRepository domain.ReservationRepository
}

func NewTrainService(
	trainRepository domain.TrainRepository,
	reservationRepository domain.ReservationRepository,
) TrainService {
	return &trainService{
		trainRepository:       trainRepository,
		reservationRepository: reservationRepository,
	}
}

func (t *trainService) GetTrainInfo(ctx context.Context, trainNumber int) (domain.Train, error) {
	trainInfo, err := t.trainRepository.FindTrainInfo(ctx, trainNumber)
	if err != nil {
		return domain.Train{}, err
	}
	return trainInfo, nil
}
