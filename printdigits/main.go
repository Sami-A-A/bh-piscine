package main

import "github.com/01-edu/z01"

func main() {
	for _, e := range "0123456789" {
		z01.PrintRune(e)
	}
	z01.PrintRune('\n')
}
