package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	a := nb
	b := RecursivePower(nb, power-1)
	nb *= b
	if nb < 0 && a > 0 {
		return 0
	} else {
		return nb
	}
}
