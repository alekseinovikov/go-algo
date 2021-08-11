package numbers

import "math"

func SplitOntoFactors(number int) []int {
	result := make([]int, 0)

	for number%2 == 0 {
		result = append(result, 2)
		number = number / 2
	}

	factor := 3
	maxFactor := int(math.Sqrt(float64(number)))
	for factor <= maxFactor {
		for number%factor == 0 {
			result = append(result, factor)
			number = number / factor

			maxFactor = int(math.Sqrt(float64(number)))
		}

		factor = factor + 2
	}

	if number > 1 {
		result = append(result, number)
	}

	return result
}
