package domain

import (
	"context"
)

type TrainRepository interface {
	FindTrainInfo(ctx context.Context, trainNumber int) (Train, error)
}
