package piscine

func IsUpper(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] < 65 || r[i] > 90 {
			return false
		}
	}
	return true
}
