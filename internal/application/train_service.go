package application

import (
	"context"

	"github.com/Mikens404/HSRSP/internal/domain"
)

type TrainService struct {
	trainRepository domain.TrainRepository
}

func NewTrainService(
	trainRepository domain.TrainRepository,
) TrainService {
	return TrainService{
		trainRepository: trainRepository,
	}
}

/* type TrainService interface {
	GetTrainInfo(ctx context.Context, trainNumber int) (domain.Train, error)
} */

func (t TrainService) GetTrainInfo(ctx context.Context, trainNumber int) (domain.Train, error) {
	trainInfo, err := t.trainRepository.GetTrainInfo(ctx, trainNumber)
	if err != nil {
		return domain.Train{}, err
	}
	return trainInfo, nil
}
