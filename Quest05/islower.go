package piscine

func IsLower(s string) bool {
	for _, e := range s {
		if e > 'z' || e < 'a' {
			return false
		}
	}
	return true
}
