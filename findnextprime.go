package student

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	for {
		isprime := true
		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {
				isprime = false
				break
			}
		}
		if isprime {
			return nb
		}
		nb++
	}
}
