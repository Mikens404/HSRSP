package domain

func CalculateDistance(boardingStation string, getOffStation string) int {
	return 1
}

func CalculateFare(distance int) int {
	switch {
	case distance <= 10:
		return 500
	case distance <= 20:
		return 600
	case distance <= 30:
		return 700
	}
	return 0
}
