package piscine

func ShoppingListSort(slice []string) []string {
	list := slice
	for i, j := 1, 0; j <= len(slice)-1; i++ {
		for _, e := range slice {
			if len(e) == i {
				list[j] = e
				j++
				break
			}
		}
	}
	return list
}
