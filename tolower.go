package piscine

func ToLower(s string) string {
	r := []rune(s)
	res := ""
	for i := range r {
		if r[i] >= 65 && r[i] <= 90 {
			res += string(r[i] + 32)
		} else {
			res += string(r[i])
		}
	}
	return res
}
