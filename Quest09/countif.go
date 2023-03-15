package piscine

func CountIf(f func(string) bool, tab []string) int {
	var count int
	for _, e := range tab {
		if f(e) {
			count++
		}
	}
	return count
}
