package piscine

func Max(a []int) int {
	var max int
	for _, e := range a {
		if e > max {
			max = e
		}
	}
	return max
}
