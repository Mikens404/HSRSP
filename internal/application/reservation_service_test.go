package application

import (
	"context"
	"errors"
	"testing"

	"github.com/Mikens404/HSRSP/internal/domain"
)

func Test_reservationService_CreateReservation(t *testing.T) {
	type fields struct {
		reservationRepository domain.ReservationRepository
	}
	type args struct {
		ctx                     context.Context
		createReservationParams domain.CreateReservationParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Repositoryからエラーが返却されない場合,エラーが返らない",
			fields: fields{
				reservationRepository: &ReservationRepositoryMock{
					InsertReservationFunc: func(ctx context.Context, createReservationParams domain.CreateReservationParams) error {
						return nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				createReservationParams: domain.CreateReservationParams{
					TrainNumber:     1,
					BoardingStation: "さんばか城",
					GetOffStation:   "一期生ハウス",
					ReservationSeats: []domain.ReservationSeat{
						{
							CarNumber:  1,
							SeatNumber: "1A",
						},
						{
							CarNumber:  1,
							SeatNumber: "1B",
						},
					},
					ReservationPeople: 2,
					CustomerInfo:      "スノーホワイトパラダイス",
				},
			},
			wantErr: false,
		},
		{
			name: "Repositoryからエラーが返却された場合,エラーが返る",
			fields: fields{
				reservationRepository: &ReservationRepositoryMock{
					InsertReservationFunc: func(ctx context.Context, createReservationParams domain.CreateReservationParams) error {
						return errors.New("")
					},
				},
			},
			args: args{
				ctx: context.Background(),
				createReservationParams: domain.CreateReservationParams{
					TrainNumber:     1,
					BoardingStation: "さんばか城",
					GetOffStation:   "一期生ハウス",
					ReservationSeats: []domain.ReservationSeat{
						{
							CarNumber:  1,
							SeatNumber: "1A",
						},
						{
							CarNumber:  1,
							SeatNumber: "1B",
						},
					},
					ReservationPeople: 2,
					CustomerInfo:      "スノーホワイトパラダイス",
				},
			},
			wantErr: true,
		},
		{
			name: "予約人数と予約席数が一致しない場合エラーが返却される",
			fields: fields{
				reservationRepository: &ReservationRepositoryMock{
					InsertReservationFunc: func(ctx context.Context, createReservationParams domain.CreateReservationParams) error {
						return nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				createReservationParams: domain.CreateReservationParams{
					TrainNumber:     1,
					BoardingStation: "さんばか城",
					GetOffStation:   "一期生ハウス",
					ReservationSeats: []domain.ReservationSeat{
						{
							CarNumber:  1,
							SeatNumber: "1A",
						},
						{
							CarNumber:  1,
							SeatNumber: "1B",
						},
					},
					ReservationPeople: 1,
					CustomerInfo:      "スノーホワイトパラダイス",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &reservationService{
				reservationRepository: tt.fields.reservationRepository,
			}
			if err := r.CreateReservation(tt.args.ctx, tt.args.createReservationParams); (err != nil) != tt.wantErr {
				t.Errorf("reservationService.CreateReservation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
