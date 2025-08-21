package presentation

import (
	"context"

	"github.com/Mikens404/HSRSP/internal/application"
)

func GetReservationSeat(ctx context.Context, params GetReservationSeatParams) (GetReservationSeatOK, error) {
	// application.GetReservationSeat が (domain.SeatReservationStatus, error) を返すと仮定
	rs, err := application.GetReservationSeat(ctx, params.TrainNumber, params.CarNumber)
	if err != nil {
		return nil, err
	}
	return rs.Reservations, nil
}
