package main

import (
	"algo/numbers"
	"fmt"
	"math"
)

func main() {
	result := numbers.TrapezoidCalculate(0.0, 5.0, 10, func(x float64) float64 {
		return 1 + x + math.Sin(2*x)
	})

	fmt.Println(result)
}
