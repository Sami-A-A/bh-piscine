package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	if n > len(a)-1 && n > 0 {
		a[n]
	} else {
		return '0'
	}
	return a[n]
}
