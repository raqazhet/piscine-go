package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 1 {
		return 1
	} else {
		raz := 1
		for i := 1; i <= nb; i++ {
			raz *= i
		}
		return raz
	}
}
