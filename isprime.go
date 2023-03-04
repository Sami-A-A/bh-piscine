package piscine

func IsPrime(nb int) bool {
	var isPrime bool
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		} else {
			isPrime = true
		}
	}
	return isPrime
}
