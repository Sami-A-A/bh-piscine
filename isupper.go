package piscine

func IsUpper(s string) bool {
	rune := []rune(s)
	counter := 0

	for i := 0; i <= len(rune)-1; i++ {
		if rune[i] >= 'A' && rune[i] <= 'Z' {
			counter++
		}
	}

	return counter == len(rune)
}
