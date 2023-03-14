package piscine

func Join(strs []string, sep string) string {
	str := ""
	for index, e := range strs {
		if index == len(strs)-1 {
			str += e
			break
		}
		str += e + sep
	}
	return str
}
