package domain

type CustomerReservationStatus int

const (
	Reserved               CustomerReservationStatus = iota // 予約済
	ProvisionalReservation                                  // 仮予約
	Cancellation                                            // キャンセル

)

type Reservation struct {
	reservationNumber     int
	reservationTrinNumber int
	reservationSeat       []SeatReservationStatus
	amount                int
	reservationPeople     int
	reservationStatus     CustomerReservationStatus
	customerInfo          string
}
