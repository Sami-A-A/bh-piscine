package piscine

func IterativePower(nb int, power int) int {
	a := nb
	for i := 1; i < power; i++ {
		nb *= a
	}
	return nb
}
