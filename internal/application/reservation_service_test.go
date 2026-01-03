package application

import (
	"context"
	"errors"
	"testing"

	"github.com/Mikens404/HSRSP/internal/domain"
)

func Test_reservationServiceImpl_CreateReservation(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		reservationRepository domain.ReservationRepository
		// Named input parameters for target function.
		createReservationParams domain.CreateReservationParams
		wantErr                 bool
	}{
		{
			name: "Repositoryからエラーが返却されない場合,エラーが返らない",
			reservationRepository: &ReservationRepositoryMock{
				InsertReservationFunc: func(ctx context.Context, createReservationParams domain.CreateReservationParams) error {
					return nil
				},
			},
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
			wantErr: false,
		},
		{
			name: "Repositoryからエラーが返却された場合,エラーが返る",
			reservationRepository: &ReservationRepositoryMock{
				InsertReservationFunc: func(ctx context.Context, createReservationParams domain.CreateReservationParams) error {
					return errors.New("repository error")
				},
			},
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
			wantErr: true,
		},
		{
			name: "予約人数と予約席数が一致しない場合エラーが返却される",
			reservationRepository: &ReservationRepositoryMock{
				InsertReservationFunc: func(ctx context.Context, createReservationParams domain.CreateReservationParams) error {
					return nil
				},
			},
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
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewReservationService(tt.reservationRepository)
			gotErr := r.CreateReservation(context.Background(), tt.createReservationParams)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateReservation() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateReservation() succeeded unexpectedly")
			}
		})
	}
}
