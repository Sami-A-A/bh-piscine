package piscine

func MakeRange(min, max int) []int {
	var a []int
	var len int = max - min
	if min < max {
		b := make([]int, len)
		for i := 0; i < len; i++ {
			b[i] = min + i
		}
		return b
	}
	return a
}
