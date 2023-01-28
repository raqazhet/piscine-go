package piscine

func LastRune(s string) rune {
	full := []rune(s)
	return full[len(full)-1]
}
