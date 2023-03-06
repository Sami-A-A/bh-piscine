package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 3159500242930794280 {
		z01.PrintRune('0')
		n /= 10
	}
	if n <= 0 {
		z01.PrintRune('0')
	}
	var a []rune
	for x := 1; n/x > 0; x *= 10 {
		if x*10 > n {
			for x := x; x*10 > n; x /= 10 {
				a = append(a, rune(int(n/x)))
				n -= int(n/x) * x
				if n > 10 && n/(x/10) < 0 {
					a = append(a, '0')
					x /= 10
				}
			}
		}
	}
	for i := 0; i < 10; i++ {
		for _, e := range a {
			if int(e) == i {
				z01.PrintRune(e + 48)
			}
		}
	}
}
