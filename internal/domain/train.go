package domain

type Train struct {
	trainNo   int
	timeTable TimeTable
	trainType int // string 絶対めんどくさい
	seatMap   SeatMap
}

type SeatMap struct { // 予約状況も一緒に格納?
}

type TimeTable struct{}
