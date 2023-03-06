package piscine

func ToLower(s string) string {
	a := []rune(s)
	for i, e := range s {
		if e >= 'A' && e <= 'Z' {
			a[i] = e + 32
		}
	}
	return string(a)
}
