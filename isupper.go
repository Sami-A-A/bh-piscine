package piscine

func IsUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 'Z' || s[i] < 'A' {
			return false
		}
	}
	return true
}
