package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	if n <= len(a) && n > 0 {
		return a[n-1]
	} else {
		return 0
	}
}
