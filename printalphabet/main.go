package main
import "github.com/01-edu/z01"

// func main() {
// 	var abc string = "abcdefghijklmnopqrstuvwxyz\n"
// 	for i := 0 ; i < 27 ; i++ {
// 		z01.PrintRune(rune(abc[i]))
// 	}
// }

func main() {
	for _, e := range "abcdefghijklmnopqrstuvwxyz/n" {
		z01.PrintRune(e)
	}
}
