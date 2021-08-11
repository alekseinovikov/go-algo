package numbers

import "math"

func FindPrimes(until int) []int {
	composites := make([]bool, until+1)
	composites[0] = true
	composites[1] = true

	for i := 4; i <= until; i = i + 2 {
		composites[i] = true
	}

	pointer := 3
	stopAt := int(math.Sqrt(float64(until)))

	for pointer <= stopAt {
		for i := pointer * 2; i <= until; i = i + pointer {
			composites[i] = true
		}

		nextPointer := pointer + 1
		for nextPointer < len(composites) && composites[nextPointer] {
			nextPointer++
		}

		pointer = nextPointer
	}

	result := make([]int, 0)
	for i := 2; i <= until; i++ {
		if !composites[i] {
			result = append(result, i)
		}
	}

	return result
}
