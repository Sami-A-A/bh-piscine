package piscine

func IsPrime(nb int) bool {
	var isPrime bool
	for i := 1; i < nb; i++ {
		if nb%i == 0 && i != 1 {
			return false
		} else {
			isPrime = true
		}
	}
	return isPrime
}
