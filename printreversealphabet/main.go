package main

import "github.com/01-edu/z01"

func main() {

	for _, e := range "zyxwvutsrqponmlkjihgfedcba" {
		z01.PrintRune(e)
	}
	z01.PrintRune('\n')
}
