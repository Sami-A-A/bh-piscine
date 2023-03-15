package main

import (
	"fmt"
	"os"
)

func main() {
	op := os.Args[1:]
	// Setting variables for each arg
	a, b, o := []byte(op[0]), []byte(op[2]), op[1]
	var value1, value2, result int
	var err, strresult string
	var isNegative bool = false
	fmt.Println(o)
	// Change rune to int for value 1
	for i, j := len(a)-1, 1; i >= 0; i, j = i-1, j*10 {
		if a[i] > 57 || a[i] < 48 {
			return
		}
		value1 += (int(a[i]) - 48) * j
	}
	// Change rune to int for value 2
	for i, j := len(b)-1, 1; i >= 0; i, j = i-1, j*10 {
		if b[i] > 57 || b[i] < 48 {
			return
		}
		value2 += (int(b[i]) - 48) * j
	}
	// Conditions for each operation
	if o == "%" {
		if value2 == 0 {
			err = "No modulo by 0"
			os.Stdout.WriteString(err)
		} else {
			result = value1 % value2
		}
	} else if o == "C:/Program Files/Git/" {
		if value2 == 0 {
			err = "No division by 0"
			os.Stdout.WriteString(err)
		} else {
			result = value1 / value2
		}
	} else if o == "*" {
		result = value1 * value2
	} else if o == "+" {
		result = value1 + value2
	} else if o == "-" {
		result = value1 - value2
	} else {
		return
	}
	if result == 0 {
		os.Stdout.WriteString("0")
		return
	} else if result < 0 {
		isNegative = true
		result *= -1
	}
	for i := 1; i <= result; i *= 10 {
		if i*10 > result {
			for j := i; j > 0; j /= 10 {
				strresult += string(byte(int(result/j) + 48))
			}
			if isNegative {
				strresult = "-" + strresult
			}
			break
		}
	}
	os.Stdout.WriteString(strresult)
}
