package piscine

func IsUpper(s string) bool {
	for _, e := range s {
		if e > 'Z' || e < 'A' {
			return false
		}
	}
	return true
}