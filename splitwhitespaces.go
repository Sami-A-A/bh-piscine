package piscine

func SplitWhiteSpaces(s string) []string {
	var arr []string
	word := ""
	for i := 0; i < len(s); i++ {
		if !(s[i] == ' ' || s[i] == '\n' || s[i] == '\t') {
			word += string(s[i])
		} else if s[i] == ' ' || s[i] == '\n' || s[i] == '\t' {
			if word == "" {
				continue
			}
			arr = append(arr, word)
			word = ""
		}
	}
	if word != "" {
		arr = append(arr, word)
	}
	return arr
}
