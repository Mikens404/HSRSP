package api

import (
	"context"

	"github.com/Mikens404/HSRSP/internal/application"
)

func GetReservationSeat(ctx context.Context, params GetReservationSeatParams) GetReservationSeatOK {
	rs := application.GetReservationSeat(ctx, params.TrainNumber, params.CarNumber)
	return GetReservationSeatOK{}
}
