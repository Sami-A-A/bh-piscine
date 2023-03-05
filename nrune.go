package piscine

func NRune(s string, n int) rune {
	ar := []rune(s)

	if n < len(ar)-1 {
		return ar[n]
	} else {
		return '0'
	}
}
