package piscine

func Index(s string, toFind string) int {
	a, b := []rune(s), []rune(toFind)
	var index, match int
	for i := 0; i < len(a); i++ {
		if a[i] == b[0] {
			index = i
			match = 0
			for j, k := 0, i; j < len(b); j, k = j+1, k+1 {
				if a[k] == b[j] {
					match++
					if match == len(b) {
						return index
					}
				} else {
					break
				}
			}
		}
	}
	return -1
}
