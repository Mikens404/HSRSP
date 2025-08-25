package domain

func CalculateDistance(boardingStation string, getOffStation string) int {
	// TODO:実装
	return 1
}

func CalculateFare(distance int) int {
	switch {
	case distance <= 0:
		return 0
	case distance <= 10:
		return 500
	case distance <= 20:
		return 600
	default:
		return 700
	}
}
