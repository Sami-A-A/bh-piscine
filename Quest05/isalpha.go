package piscine

func IsAlpha(s string) bool {
	for _, e := range s {
		if (e < 'a' || e > 'z') && (e < 'A' || e > 'Z') && (e < 48 || e > 57) {
			return false
		}
	}
	return true
}
