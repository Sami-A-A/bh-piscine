package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	y := n - 1
	if y < len(a)-1 && y >= 0 {
		return a[y]
	} else {
		return 0
	}
}
