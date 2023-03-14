package piscine

func ShoppingListSort(slice []string) []string {
	list := make([]string, len(slice))
	for i, j := 1, 0; j < len(slice); i, j = i+1, j+1 {
		for _, e := range slice {
			if len(e) == i {
				list[j] = e
				break
			}
		}
	}
	return list
}
