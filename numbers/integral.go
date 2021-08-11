package numbers

func RectangleCalculate(from float64, to float64, intervals int, f func(float64) float64) float64 {
	dx := (to - from) / float64(intervals)

	result := 0.0
	x := from

	for i := 1; i <= intervals; i++ {
		area := dx * f(x)
		result += area
		x += dx
	}

	return result
}

func TrapezoidCalculate(from float64, to float64, intervals int, f func(float64) float64) float64 {
	dx := (to - from) / float64(intervals)

	result := 0.0
	x := from

	for i := 1; i <= intervals; i++ {
		area := (dx * (f(x) + f(x+dx))) / 2
		result += area
		x += dx
	}

	return result
}
