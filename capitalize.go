package piscine

func Capitalize(s string) string {
	r := []rune(s)
	h := true
	for i := 0; i < len(r); i++ {
		if r[i] >= 97 && r[i] <= 122 || r[i] >= 65 && r[i] <= 90 || r[i] >= 48 && r[i] <= 57 {
			if h && r[i] >= 97 && r[i] <= 122 {
				r[i] = (r[i] - 32)
			} else if r[i] >= 65 && r[i] <= 90 && !h {
				r[i] = (r[i] + 32)
			}
			h = false
		} else {
			h = true
		}
	}
	return string(r)
}
