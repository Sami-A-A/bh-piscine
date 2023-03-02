package piscine

import (
	"math"
)

func BasicAtoi(s string) int {
	var num int
	var r string
	// reverse string
	for _, char := range s {
		r = string(char) + r
	}
	// convert to rune array
	runes := []rune(r)
	// loop over array to calculate string value based on rune values
	for i := 0; i < len(runes); i++ {
		if (int(runes[i]) - 48) == 0 {
			continue
		} else {
			num = (int(runes[i]) - 48) * int(math.Pow(10, float64(i)))
		}
	}
	return num
}
