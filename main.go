package main

import (
	"algo/randomization"
	"fmt"
)

func main() {
	size := 1000_000
	arr := make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = i
	}

	randomization.RandomizeIntArray(arr)

	fmt.Println(arr)
}
