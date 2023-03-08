package piscine

func Split(s, sep string) []string {
	var str_slc []string
	var word string
	var str_check bool
	for x := 0; x < len(s); x++ {
		for y, x2 := 0, x; y < len(sep); y, x2 = y+1, x2+1 {
			if s[x2] == sep[y] {
				str_check = true
			} else {
				str_check = false
				break
			}
		}
		if !str_check {
			word += string(s[x])
		} else if str_check {
			str_slc = append(str_slc, word)
			word, x = "", x+len(sep)-1
		}
	}
	return str_slc
}
