package piscine

func SplitWhiteSpaces(s string) []string {
	var a []string
	var b string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' || s[i] == '\t' {
			a = append(a, b)
			b = ""
		} else if !(s[i] == ' ' || s[i] == '\n' || s[i] == '\t') {
			b += string(s[i])
		}
	}
	if b != "" {
		a = append(a, b)
	}
	return a
}
