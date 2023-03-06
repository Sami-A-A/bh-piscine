package piscine

func ToUpper(s string) string {
	alr := []rune(s)
	res := []rune{}

	for j := 0; j <= len(alr)-1; j++ {
		for k := 'A'; k <= 'Z'; k++ {
			if alr[j] >= 'a' && alr[j] <= 'z' {
				res = append(res, alr[j]-32)
				break
			} else {
				res = append(res, alr[j])
				break
			}
		}
	}
	return string(res)
}
