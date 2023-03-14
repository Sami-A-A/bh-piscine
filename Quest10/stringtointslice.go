package piscine

func StringToIntSlice(str string) []int {
	runeSlice := []rune(str)
	intSlice := make([]int, len(runeSlice))
	for i := 0; i < len(runeSlice)-1; i++ {
		intSlice[i] = int(runeSlice[i])
	}
	return intSlice
}
