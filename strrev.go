package piscine

func StrRev(s string) string {
	raz := []rune(s)
	word := ""
	last := len(raz) - 1
	for i := last; i >= 0; i-- {
		word += string(raz[i])
	}
	return word
}

