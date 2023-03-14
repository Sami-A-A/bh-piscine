package piscine

func ShoppingListSort(slice []string) []string {
	list := []string{}
	for i := 1; len(list) < len(slice); i++ {
		for _, e := range slice {
			if len(e) == i {
				list = append(list, e)
			}
		}
	}
	return list
}
