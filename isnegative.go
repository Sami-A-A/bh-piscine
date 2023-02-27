package main

import "github.com/01-edu/z01"

func isNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
}

func main() {
	isNegative(-1)
}
