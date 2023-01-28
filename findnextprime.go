package piscine

func P(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i < (nb/2)+1; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	for {
		if P(nb) == true {
			return nb
		}
		nb++
	}
}
