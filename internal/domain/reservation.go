package domain

type CustomerReservationStatus int

const (
	Reserved               CustomerReservationStatus = iota // 予約済
	ProvisionalReservation                                  // 仮予約
	Cancellation                                            // キャンセル

)

type Reservation struct {
	reservationNumber     int
    reservationTrainNumber int
	reservationSeat       []TrainReservationStatus
	amount                int
	reservationPeople     int
	reservationStatus     CustomerReservationStatus
    customerInfo          string
}
