package piscine

func PrintNbr(nbr int) {
}

func ForEach(f func(int), a []int) {
	for _, e := range a {
		PrintNbr(e)
	}
}
