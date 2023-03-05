package piscine

func AlphaCount(s string) int {
	a := []rune(s)
	count := 0
	for i := 0; i < len(s); i++ {
		if (a[i] >= 'a' && a[i] <= 'z') || (a[i] >= 'A' && a[i] <= 'Z') {
			count++
		}
	}
	return count
}
