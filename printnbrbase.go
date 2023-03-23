package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	bases := []rune(base)
	for i := 0; i < len(bases); i++ {
		for j := i + 1; j < len(bases); j++ {
			if bases[i] == bases[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	if nbr < 0 {
		nbr *= -1
		z01.PrintRune('-')
	}
	for n := 1; nbr > n; n *= len(bases) {
		if n*len(bases) > nbr {
			for n > 0 {
				z01.PrintRune(bases[int(nbr/n)])
				nbr %= n
				n /= len(bases)
			}
			break
		}
	}
}
