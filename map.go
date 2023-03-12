package piscine

func Map(f func(int) bool, a []int) []bool {
	m := make([]bool, len(a))
	for i, e := range a {
		m[i] = f(e)
	}
	return m
}
