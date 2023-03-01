package piscine

func DivMod(a int, b int, div *int, mod *int) {
	dividend := a / b
	remainder := a % b
	*div = dividend
	*mod = remainder
}
