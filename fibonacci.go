package piscine

var a, b, c = 0, 1, 0

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	}
	c = a + b
	a, b = b, c
	Fibonacci(index - 1)
	return a
}
