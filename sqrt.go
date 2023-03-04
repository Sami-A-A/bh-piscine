package piscine

import "math"

func Sqrt(nb int) int {
	sr := math.Sqrt(float64(nb))
	if int(sr)*int(sr) == nb {
		return int(sr)
	} else {
		return 0
	}
}
