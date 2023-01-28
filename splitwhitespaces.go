package piscine

func SplitWhiteSpaces(s string) []string {
	var raz []string
	for i := 0; i < len(s); i++ {
		w := ""
		for s[i] != ' ' && s[i] != '\n' && s[i] != 9 {
			w += string(s[i])
			if i == len(s)-1 {
				break
			}
			i++
		}
		if w != "" {
			raz = append(raz, w)
		}
	}
	return raz
}
