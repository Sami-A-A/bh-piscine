package piscine

func IsPrime(nb int) bool {
	var isPrime bool
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			isPrime = true
			return isPrime
		} else {
			isPrime = false
		}
	}
	return isPrime
}
