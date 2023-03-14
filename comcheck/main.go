package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, e := range args {
		if e == "01" || e == "galaxy" || e == "galaxy 1101" {
			fmt.Println("Alert!!!")
		}
	}
}
