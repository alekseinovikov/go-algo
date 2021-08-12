package main

import (
	"fmt"
)

func validUtf8(data []int) bool {
	waiting := 0
	for _, byte := range data {
		str := fmt.Sprintf("%08b", byte)

		if waiting > 0 {
			if str[0] == '0' {
				return false
			} else if str[:2] == "10" {
				waiting--
			} else if str[:2] == "11" {
				return false
			} else {
				return false
			}
		} else if waiting <= 0 {
			if str[:2] == "10" {
				return false
			} else if str[:3] == "110" {
				waiting = 1
			} else if str[:4] == "1110" {
				waiting = 2
			} else if str[:5] == "11110" {
				waiting = 3
			} else if str[:1] == "0" {
				continue
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return waiting <= 0
}

func main() {
	result := validUtf8([]int{237})

	fmt.Printf("Result: %t Expected: %t ", result, false)
}
