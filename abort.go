package piscine

func Abort(a, b, c, d, e int) int {
	nrange := []int{a, b, c, d, e}
	var median, count int
	for _, n := range nrange {
		count = 0
		for _, e := range nrange {
			if n > e {
				count += 1
			}
		}
		if count == 2 {
			median = n
		}
	}
	return median
}
