package piscine

func FindNextPrime(nb int) int {
	if nb > 1 {
		for a := nb; a >= nb; a++ {
			for b := 1; b <= a; b++ {
				if a%b == 0 && b != 1 && b != a {
					break
				} else if b == a {
					return a
				}
			}
		}
	}
	return -1
}
