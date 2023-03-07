package piscine

func BasicJoin(elems []string) string {
	var a string
	for _, e := range elems {
		a += e
	}
	return a
}
