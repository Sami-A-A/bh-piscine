package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
	rns := []rune{'x', ' ', '=', ' ', '4', '1', ',', ' ', 'y', ' ', '=', ' ', '2', '1'}
	for _, e := range rns {
		z01.PrintRune(e)
	}
}

func main() {
	points := &point{}
	setPoint(points)
}
