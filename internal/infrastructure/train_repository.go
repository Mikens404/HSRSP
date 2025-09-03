package infrastructure

import (
	"context"
	"errors"
	"time"

	"github.com/Mikens404/HSRSP/internal/domain"
)

// いつかDB設計する😭

type trainRepositoryImpl struct{}

func NewTrainRepository() trainRepositoryImpl {
	return trainRepositoryImpl{}
}

func (t *trainRepositoryImpl) FindTrainInfo(ctx context.Context, trainNumber int) (domain.Train, error) {
	if trainNumber != 1 {
		return domain.Train{}, errors.New("指定された列車は存在しません")
	}
	timeTable := domain.TimeTable{
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
		TimeTable:   timeTable,
		TrainType:   domain.Rapid,
	}
	return responseTrainInfo, nil
}
