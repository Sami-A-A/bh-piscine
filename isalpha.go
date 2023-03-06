package piscine

func IsAlpha(s string) bool {
	for _, e := range s {
		if (e < 'a' || e > 'z') && (e < 'A' || e > 'Z') {
			return false
		}
	}
	return true
}
