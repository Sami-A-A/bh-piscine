package piscine

func IsPrintable(s string) bool {
	for _, e := range s {
		if e < 14 {
			return false
		}
	}
	return true
}
