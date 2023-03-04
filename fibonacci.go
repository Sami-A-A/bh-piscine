package piscine

var a, b = 0, 1

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	}
	a, b = b, a+b
	Fibonacci(index - 1)
	return a
}
