package piscine

func Join(strs []string, sep string) string {
	raz := ""
	for i := 0; i < len(strs); i++ {
		if i < len(strs)-1 {
			raz += strs[i] + sep
		} else {
			raz += strs[i]
		}
	}
	return raz
}
