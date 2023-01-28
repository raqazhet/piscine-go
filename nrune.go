package piscine

func NRune(s string, n int) rune {
	raz := []rune(s)
	len := 0
	for index := range raz {
		len = index
	}
	if n-1 >= 0 && n-1 <= len {
		return (raz[n-1])
	}
	return 0
}