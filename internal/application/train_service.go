package application

import (
	"context"

	"github.com/Mikens404/HSRSP/internal/domain"
	"github.com/Mikens404/HSRSP/internal/infrastructure"
)

type TrainService interface {
	GetTrainInfo(ctx context.Context, tarinNumber int) (domain.Train, error)
}

func GetTrainInfo(ctx context.Context, tarinNumber int) (domain.Train, error) {
	trianInfo, err := infrastructure.FindTrainInfo(ctx, 1)
	if err != nil {
		return domain.Train{}, err
	}
	return trianInfo, nil
}
