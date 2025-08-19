package domain

type CustomerReservationStatus int

const (
	Reserved               CustomerReservationStatus = iota // 予約済
	ProvisionalReservation                                  // 仮予約
	Cancellation                                            // キャンセル

)

type Reservation struct {
	reservationNumber      int
	reservationTrainNumber int
	reservationSeat        []ReservationSeat
	boardingStation        string
	get_offStation         string
	amount                 int
	reservationPeople      int
	reservationStatus      CustomerReservationStatus
	customerInfo           string
}

type ReservationSeat struct {
	carNumber    int    // 何両目
	reservations string // 席番
}
