package numbers

func Power(number int64, power int64) int64 {
	currValue := number
	currPower := int64(1)
	nextPower := currPower * 2
	for nextPower <= power {
		currValue = currValue * currValue

		currPower = nextPower
		nextPower = nextPower * 2
	}

	restPower := power - currPower

	if restPower <= 0 {
		return currValue
	}

	restValue := int64(1)

	for i := int64(0); i < restPower; i++ {
		restValue = restValue * number
	}

	return currValue * restValue
}
