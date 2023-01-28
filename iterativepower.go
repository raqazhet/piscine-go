package piscine

func IterativePower(nb int, power int) int {
	raz := 1
	if power < 0 {
		raz = 0
	} else {
		for i := 1; i <= power; i++ {
			raz *= nb
		}
	}
	return raz
}
