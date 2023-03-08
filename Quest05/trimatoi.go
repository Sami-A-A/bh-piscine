package piscine

func TrimAtoi(s string) int {
	var a []rune
	var b int = 1
	var result int
	isPositive := true
	for _, e := range s {
		if e == '-' && len(a) == 0 {
			isPositive = false
		}
		if e >= 48 && e <= 57 {
			a = append(a, e)
		}
	}
	for i := 0; i < len(a)-1; i++ {
		b *= 10
	}
	for i := 0; i < len(a); i, b = i+1, b/10 {
		result += int(int(a[i]-48) * b)
	}
	if !isPositive {
		return result * -1
	} else {
		return result
	}
}
