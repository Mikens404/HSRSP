package infrastructure

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Mikens404/HSRSP/internal/domain"
)

func getMockTrainData() domain.Train {
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
	return domain.Train{
		TrainNumber: 1,
		TimeTable:   timeTable,
		TrainType:   domain.Rapid}
}

func Test_trainRepositoryImpl_FindTrainInfo(t *testing.T) {
	type args struct {
		ctx         context.Context
		trainNumber int
	}
	tests := []struct {
		name    string
		tr      trainRepositoryImpl
		args    args
		want    domain.Train
		wantErr bool
	}{
		{
			name: "ハードコーディングされた情報を取得できる",
			args: args{
				ctx:         context.Background(),
				trainNumber: 1,
			},
			want:    getMockTrainData(),
			wantErr: false,
		},
		{
			name: "存在しない列車番号を指定したとき,エラーを返却できる",
			args: args{
				ctx:         context.Background(),
				trainNumber: 2,
			},
			want:    domain.Train{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := trainRepositoryImpl{}
			got, err := tr.FindTrainInfo(tt.args.ctx, tt.args.trainNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("trainRepositoryImpl.FindTrainInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trainRepositoryImpl.FindTrainInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
