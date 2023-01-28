package piscine

func IsNumeric(s string) bool {
	r := []rune(s)
	for i := range r {
		if r[i] >= 47 && r[i] <= 58 {
		} else {
			return false
		}
	}
	return true
}
