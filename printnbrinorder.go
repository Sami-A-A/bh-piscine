package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	var a []rune
	for x := 1; n > 0; x *= 10 {
		if n < 10 {
			a = append(a, rune(n))
			n = 0
		} else if x*10 > n {
			a = append(a, rune(int(n/x)))
			n -= int(n/x) * x
			x = 1
		}
	}
	for i := 0; i < 10; i++ {
		for _, e := range a {
			if int(e) == i {
				z01.PrintRune(e + 48)
				break
			}
		}
	}
}
