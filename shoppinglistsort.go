package piscine

func ShoppingListSort(slice []string) []string {
	for index, length := 0, 0; index < len(slice); index++ {
		for i, e := range slice {
			if len(e) == length {
				slice[i], slice[index] = e, slice[i]
				length++
				break
			}
		}
	}
	return slice
}
