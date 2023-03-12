package piscine

import "github.com/01-edu/z01"

func PrintNbr(nbr int) {
	z01.PrintRune(rune(nbr))
}

func ForEach(f func(int), a []int) {
	for _, e := range a {
		PrintNbr(e)
	}
}
