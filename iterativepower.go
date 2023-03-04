package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	a := nb
	for i := 1; i < power; i++ {
		nb *= a
	}
	if nb < 0 && a > 0 {
		return 0
	} else {
		return nb
	}
}
