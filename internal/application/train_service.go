package application

import (
	"context"

	"github.com/Mikens404/HSRSP/internal/domain"
	"github.com/Mikens404/HSRSP/internal/infrastructure"
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
	trainInfo, err := infrastructure.FindTrainInfo(ctx, trainNumber)
	if err != nil {
		return domain.Train{}, err
	}
	return trainInfo, nil
}
