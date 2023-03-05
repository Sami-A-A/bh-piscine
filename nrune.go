package piscine

func NRune(s string, n int) rune {
	ar := []rune(s)
	if n < len(ar) {
		return ar[n-1]
	} else {
		return '0'
	}
}
