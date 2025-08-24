package domain

type CustomerReservationStatus int

const (
	Reserved               CustomerReservationStatus = iota // 予約済
	ProvisionalReservation                                  // 仮予約
	Cancellation                                            // キャンセル

)

// 予約情報
type Reservation struct {
	ReservationNumber int
	TrainNumber       int
	BoardingStation   string
	GetOffStation     string
	ReservationSeat   []ReservationSeat
	Amount            int
	ReservationPeople int
	ReservationStatus CustomerReservationStatus
	CustomerInfo      string
}

type ReservationSeat struct {
	CarNumber  int    // 何両目
	SeatNumber string // 席番
}

// 号車ごとの予約情報
type SeatReservationStatus struct {
	CarNumber    int             // 何両目
	Reservations map[string]bool // [席番]
}

type CreateReservationParams struct {
	TrainNumber       int
	BoardingStation   string
	GetOffStation     string
	ReservationSeats  []ReservationSeat
	ReservationPeople int
	CustomerInfo      string
}

type UpdateReservationParams struct {
	ReservationNumber int
	TrainNumber       int
	BoardingStation   string
	GetOffStation     string
	ReservationSeats  []ReservationSeat
	ReservationPeople int
	ReservationStatus CustomerReservationStatus
	CustomerInfo      string
}
