package piscine

func MakeRange(min, max int) []int {
	var a []int
	if min < max {
		b := make([]int, max-min)
		for i := 0; i < len(b); i++ {
			b[i] = min + i
		}
		return b
	}
	return a
}
