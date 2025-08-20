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
	TrainNumber int
	TimeTable   TimeTable
	TrainType   TrainType
}

type TimeTable []StopStationList

type StopStationList struct {
	StationName   string
	ArrivalTime   time.Time
	DepartureTime time.Time
}
