package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	itembytes := []byte(str)
	list := make(map[string]int)
	var item string
	for _, e := range itembytes {
		if (e >= 'a' && e <= 'z') || (e >= 'A' && e <= 'Z') {
			item += string(byte(e))
		} else if e == 32 {
			list[item] += 1
			item = ""
		}
	}
	list[item] += 1
	return list
}
