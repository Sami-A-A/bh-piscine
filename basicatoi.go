package piscine

func BasicAtoi(s string) int {
	var num int
	var r string
	// reverse string
	for _, char := range s {
		r = string(char) + r
	}
	// convert to rune array
	runes := []rune(r)
	// loop over array to calculate string value based on rune values
	for i := 0; i < len(runes); i++ {
		if (int(runes[i]) - 48) == 0 {
			continue
		} else {
			// create exponent function
			var power int = 10
			for j := 0; j < i; j++ {
				power *= 10
			}
			num = num + (int(runes[i])-48)*(power/10)
		}
	}
	return num
}
