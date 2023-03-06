package piscine

func ToUpper(s string) string {
	a := []rune(s)
	for i, e := range s {
		if e >= 'a' && e <= 'z' {
			a[i] = e - 32
		}
	}
	return string(a)
}
