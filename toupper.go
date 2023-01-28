package piscine

func ToUpper(s string) string {
	r := []rune(s)
	res := ""
	for i := range r {
		if r[i] >= 97 && r[i] <= 122 {
			res += string(r[i] - 32)
		} else {
			res += string(r[i])
		}
	}
	return res
}
