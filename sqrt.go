package piscine

func Sqrt(nb int) int {
	i := 0
	for i < nb {
		i++
		if i*i == nb {
			return i
		}
	}
	return 0
}
