package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args[1:]) > 1 {
		fmt.Println("Too many arguments")
		return
	} else if len(os.Args[1:]) < 1 {
		fmt.Println("File name missing")
		return
	}
	file, err := os.Open("quest8.txt")
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	arr := make([]byte, 14)
	file.Read(arr)
	fmt.Println(string(arr))
	file.Close()
}
