package piscine

func IsNumeric(s string) bool {
	for _, e := range s {
		if e < 48 || e > 57 {
			return false
		}
	}
	return true
}
