package infrastructure

import (
	"context"
	"errors"
	"time"

	"github.com/Mikens404/HSRSP/internal/domain"
)

// いつかDB設計する😭

func FindTrainInfo(ctx context.Context, tarinNumber int) (domain.Train, error) {
	if tarinNumber != 1 {
		return domain.Train{}, errors.New("指定された列車は存在しません")
	}
	timeTbale := domain.TimeTable{
		domain.StopStationList{
			StationName:   "さんばか城",
			ArrivalTime:   time.Date(2019, 03, 22, 16, 00, 00, 00, time.UTC),
			DepartureTime: time.Date(2019, 03, 22, 16, 00, 30, 00, time.UTC),
		},
		domain.StopStationList{
			StationName:   "いぬいキングダム",
			ArrivalTime:   time.Date(2019, 03, 22, 16, 10, 00, 00, time.UTC),
			DepartureTime: time.Date(2019, 03, 22, 16, 11, 00, 00, time.UTC),
		},
		domain.StopStationList{
			StationName:   "一期生ハウス",
			ArrivalTime:   time.Date(2019, 03, 22, 16, 18, 00, 00, time.UTC),
			DepartureTime: time.Date(2019, 03, 22, 16, 18, 30, 00, time.UTC),
		},
	}
	responseTrainInfo := domain.Train{
		TrainNumber: 1,
		TimeTable:   timeTbale,
		TrainType:   1,
	}
	return responseTrainInfo, nil
}
