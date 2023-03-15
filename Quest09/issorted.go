package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	isSortedAscendingly := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i+1], a[i]) < 0 && isSortedAscendingly {
			isSortedAscendingly = false
		} else if f(a[i+1], a[i]) > 0 && !isSortedAscendingly {
			return false
		}
	}
	return true
}
