package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, e := range a {
		runes := []rune(e)
		for i := 0; i < len(e); i++ {
			z01.PrintRune(runes[i])
		}
		z01.PrintRune('\n')
	}
}
