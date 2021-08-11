package numbers

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomizeArray(arr []interface{}) {
	arrayLength := len(arr)
	if arrayLength <= 1 {
		return
	}

	for i := 0; i < arrayLength; i++ {
		randomIndex := rand.Intn(arrayLength + 1)
		buff := arr[i]
		arr[i] = arr[randomIndex]
		arr[randomIndex] = buff
	}
}

func RandomizeIntArray(arr []int) {
	arrayLength := len(arr)
	if arrayLength <= 1 {
		return
	}

	rand.Shuffle(len(arr), func(i, j int) {
		buff := arr[i]
		arr[i] = arr[j]
		arr[j] = buff
	})
}
