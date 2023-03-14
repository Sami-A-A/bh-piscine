package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	itembytes := []byte(str)
	list := make(map[string]int)
	var item string
	for _, e := range itembytes {
		if e == 32 && item != "" {
			list[item] += 1
			item = ""
		} else if e != 32 {
			item += string(byte(e))
		}
	}
	list[item] += 1
	return list
}
