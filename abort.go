package piscine

func Abort(a, b, c, d, e int) int {
	raz := []int{a, b, c, d, e}
	for r := 0; r < len(raz)-1; r++ {
		for q := r + 1; q < len(raz); q++ {
			min := raz[r]
			if raz[q] < min {
				raz[r], raz[q] = raz[q], raz[r]
			}
		}
	}
	return raz[2]
}
