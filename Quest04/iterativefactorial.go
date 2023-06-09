package piscine

func IterativeFactorial(nb int) int {
	if nb > 0 && nb < 24 {
		for i := nb - 1; i > 1; i-- {
			nb = nb * i
		}
		return nb
	} else if nb == 0 {
		return 1
	} else {
		return 0
	}
}
