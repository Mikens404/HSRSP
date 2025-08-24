package infrastructure

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Mikens404/HSRSP/internal/domain"
)

func GetMockDate() domain.Train {
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
	return domain.Train{TrainNumber: 1,
		TimeTable: timeTbale,
		TrainType: 1}
}

func TestFindTrainInfo(t *testing.T) {
	type args struct {
		ctx         context.Context
		tarinNumber int
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Train
		wantErr bool
	}{
		{
			name: "ハードコーディングされた情報を取得できる",
			args: args{
				ctx:         context.Background(),
				tarinNumber: 1,
			},
			want:    GetMockDate(),
			wantErr: false,
		},
		{
			name: "存在しない列車番号を指定したとき,エラーを返却できる",
			args: args{
				ctx:         context.Background(),
				tarinNumber: 2,
			},
			want:    domain.Train{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindTrainInfo(tt.args.ctx, tt.args.tarinNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindTrainInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindTrainInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
