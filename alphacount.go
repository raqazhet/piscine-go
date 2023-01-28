package piscine

func AlphaCount(s string) int {
	a := []rune(s)
	count := 0
	for i := 0; i < len(a); i++ {
		if int(a[i]) >= 65 && int(a[i]) <= 90 || int(a[i]) >= 97 && int(a[i]) <= 122 {
			count++
		}
	}
	return count
}
