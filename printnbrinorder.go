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
			break
		} else if x*10 > n {
			a = append(a, rune(int(n/x)))
			n -= int(n/x) * x
			x = 1
		}
		if x > n {
			break
		}
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if int(a[j]) == i {
				z01.PrintRune(a[j] + 48)
				break
			}
		}
		if i == 10 {
			break
		}
	}
}
