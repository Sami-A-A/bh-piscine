package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	b := []rune(toFind)
	var index int
	for i := 0; i < len(a); i++ {
		if len(b) > 1 {
			for j := 0; j < len(b); j++{
				if a[i] == b[j] {
					continue
				} else {
					return -1
				}
			}
		} else if a[i] == b[0] {
			index = i
		}
	}
	return index
}