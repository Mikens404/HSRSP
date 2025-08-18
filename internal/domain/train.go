package domain

import "time"

type TrainType int

const (
	Local          TrainType = iota // 普通
	Rapid                           // 快速
	Express                         // 急行
	LimitedExpress                  // 特急
)

type Train struct {
	trainNo           int
	timeTable         TimeTable
	trainType         TrainType
	reservationStatus []ReservationStatus
}

type ReservationStatus struct {
	carNumber    int             // 何両目
	reservations map[string]bool // [席番]
}

type TimeTable []StationStopList

type StationStopList struct {
	StationName   string
	ArrivalTime   time.Time
	DepartureTime time.Time
}
