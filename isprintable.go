package piscine

func IsPrintable(s string) bool {
	r := []rune(s)
	for i := range r {
		if r[i] >= 32 && r[i] <= 126 {
		} else {
			return false
		}
	}
	return true
}
