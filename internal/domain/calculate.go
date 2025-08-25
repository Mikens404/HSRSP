package domain

func CalculateDistance(boardingStation string, getOffStation string) int {
	return 1
}

func CalculateFare(distance int) int {
	var responseAmount int
	switch {
	case distance <= 0:
		responseAmount = 0
	case distance <= 10:
		responseAmount = 500
	case distance <= 20:
		responseAmount = 600
	case distance <= 30:
		responseAmount = 700
	default:
		responseAmount = 700
	}
	return responseAmount
}
