package piscine

func Swap(a *int, b *int) {
	ap := *a
	bp := *b
	*a = bp
	*b = ap
}
