package domain

import (
	"context"
)

type TrainRepository interface {
	GetTrainInfo(ctx context.Context, tarinNumber int) (Train, error)
}
