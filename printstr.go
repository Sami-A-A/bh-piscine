package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, e := range s {
		z01.PrintRune(e)
	}
	z01.PrintRune('\n')
}
