package presentation

import (
	"context"

	"github.com/Mikens404/HSRSP/internal/application"
)

type handler struct {
	reservationService application.ReservationService
	trainService       application.TrainService
}

func NewHandler(
	reservationService application.ReservationService,
	trainService application.TrainService,

) handler {
	return handler{
		reservationService: reservationService,
		trainService:       trainService,
	}
}

// 号車ごとの予約状況
func (h handler) GetReservationSeat(ctx context.Context, params GetReservationSeatParams) (GetReservationSeatOK, error) {
	rs, err := h.reservationService.GetReservationSeat(ctx, params.TrainNumber, params.CarNumber)
	if err != nil {
		return nil, err
	}
	return rs.Reservations, nil
}
