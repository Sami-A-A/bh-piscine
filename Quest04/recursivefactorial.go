package piscine

func RecursiveFactorial(nb int) int {
	if nb > 0 && nb < 24 {
		a := nb - 1
		b := RecursiveFactorial(a)
		nb *= b
	} else if nb == 0 {
		return 1
	} else {
		return 0
	}
	return nb
}
