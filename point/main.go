package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	rns := []rune{52, 50, 50, 49}
	str := "x = v, y = w"
	for _, e := range str {
		if e == 'v' || e == 'w' {
			for i, r := range rns {
				if e == 'v' && i > 1 {
					break
				}
				if e == 'w' && i <= 1 {
					continue
				}
				z01.PrintRune(r)
			}
			continue
		}
		z01.PrintRune(e)
	}
	z01.PrintRune('\n')
}
