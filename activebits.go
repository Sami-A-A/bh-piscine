package piscine

func ActiveBits(n int) int {
	var active int
	for i, j := 0, 128; j >= 1; i, j = i+1, j/2 {
		if n >= j {
			n -= j
			active += 1
		}
	}
	return active
}
