package piscine

func SplitWhiteSpaces(s string) []string {
	var a []string
	var b string
	for i := 0; i < len(s); i++ {
		if rune(s[i]) <= 32 {
			a = append(a, b)
			b = ""
		} else if s[i] > 32 {
			b += string(s[i])
		}
	}
	a = append(a, b)
	return a
}
