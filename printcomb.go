package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	digits := "0123456789"
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				if a < b && b < c {
					z01.PrintRune(rune(digits[a]))
					z01.PrintRune(rune(digits[b]))
					z01.PrintRune(rune(digits[c]))
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
