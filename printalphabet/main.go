package main
import "github.com/01-edu/z01"

func main() {
	for _,e := range "abcdefghijklmnopqrstuvwxyz\n"{
		z01.PrintRune(e)
	}
}
