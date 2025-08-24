package domain

import (
	"context"
)

type TrainRepository interface {
	GetTrainInfo(ctx context.Context, trainNumber int) (Train, error)
}
