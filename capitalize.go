package piscine

func Capitalize(s string) string {
	a := []rune(s)
	initial := true
	for i, e := range a {
		if (e >= 'a' && e <= 'z') && initial {
			a[i] = e - 32
			initial = false
		} else if (a[i] >= 'A' && a[i] <= 'Z' || (a[i] >= 48 && a[i] <= 57)) && initial {
			initial = false
		} else if a[i] >= 'A' && a[i] <= 'Z' && !initial {
			a[i] = e + 32
		} else if (a[i] < 'a' || a[i] > 'z') && (a[i] < 'A' || a[i] > 'Z') && (a[i] < 48 || a[i] > 57) && !initial {
			initial = true
		}
	}
	s = string(a)
	return s
}
