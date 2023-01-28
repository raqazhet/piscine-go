package piscine

func IsLower(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] < 97 || r[i] > 122 {
			return false
		}
	}
	return true
}
