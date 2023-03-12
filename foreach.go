package piscine

func PrintNbr(nbr int) {
	print(nbr)
}

func ForEach(f func(int), a []int) {
	for _, e := range a {
		PrintNbr(e)
	}
}
