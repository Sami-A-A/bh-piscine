package piscine

func Join(strs []string, sep string) string {
	var a string
	for i, e := range strs {
		if i == len(strs)-1 {
			a += e
		} else {
			a += e + sep
		}
	}
	return a
}
