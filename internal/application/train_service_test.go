package application

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Mikens404/HSRSP/internal/domain"
)

func Test_trainService_GetTrainInfo(t *testing.T) {
	type fields struct {
		trainRepository domain.TrainRepository
	}
	type args struct {
		ctx         context.Context
		trainNumber int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Train
		wantErr bool
	}{
		{
			name: "Repositoryからデータが取得出来る",
			fields: fields{
				trainRepository: &TrainRepositoryMock{
					FindTrainInfoFunc: func(ctx context.Context, trainNumber int) (domain.Train, error) {
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
					},
				},
			},
			args: args{
				ctx:         context.Background(),
				trainNumber: 1,
			},
			want: domain.Train{
				TrainNumber: 1,
				TimeTable: domain.TimeTable{
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
				},
				TrainType: domain.Rapid,
			},
			wantErr: false,
		},
		{
			name: "Repositoryからエラーが返却されたとき,エラーが返る",
			fields: fields{
				trainRepository: &TrainRepositoryMock{
					FindTrainInfoFunc: func(ctx context.Context, trainNumber int) (domain.Train, error) {
						return domain.Train{}, errors.New("")
					},
				},
			},
			args: args{
				ctx:         context.Background(),
				trainNumber: 1,
			},
			want:    domain.Train{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &trainService{
				trainRepository: tt.fields.trainRepository,
			}
			got, err := tr.GetTrainInfo(tt.args.ctx, tt.args.trainNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("trainService.GetTrainInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trainService.GetTrainInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
