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
	arr := []string{"x = ", string(x1), string(x2), ", y = ", string(y1), string(y2)}
	for _, s := range arr {
		for _, e := range s {
			z01.PrintRune(e)
		}
	}
	z01.PrintRune('\n')
}
