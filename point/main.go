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
	x1, x2, y1, y2 := points.x+10, points.x+8, points.y+29, points.y+28
	rns := []int{x1, x2, y1, y2}
	str := "x = v, y = v"
	for _, e := range str {
		if e == 'v' {
			for _, r := range rns {
				z01.PrintRune(rune(r))
			}
			continue
		}
		z01.PrintRune(e)
	}
	z01.PrintRune('\n')
}
