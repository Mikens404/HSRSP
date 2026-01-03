package domain

import "errors"

var (
	// 指定された列車が存在しない場合
	ErrTrainNotFound = errors.New("The specified train does not exist.")
)

var (
	// 予約人数と予約座席が一致しない場合
	ErrNotmatchReservationNumber = errors.New("The number of reservations does not match.")
)
