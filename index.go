package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	b := []rune(toFind)
	var index, match int
	for i := 0; i < len(a); i++ {
		if a[i] == b[0] {
			index = i
			for j := 0; j < len(b); i, j = i+1, j+1 {
				if a[i] == b[j] {
					match++
					if match == len(b) {
						return index
					} else {
						continue
					}
				}
			}
		} else if a[i] != b[0] && i-1 == len(a) {
			return -1
		}
	}
	return index
}
