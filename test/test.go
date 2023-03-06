package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(piscine.Index("Salut!", "alut!"))
	fmt.Println(piscine.Index("asdfasdf!", "asdf!"))
}
