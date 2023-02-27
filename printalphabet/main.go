package main
import "github.com/01-edu/z01"

func main() {
	var abc string = "abcdefghijklmnopqrstuvwxyz"
	for i := 0 ; i < 26 ; i++ {
		z01.PrintRune(rune(abc[i]))
	}
}
