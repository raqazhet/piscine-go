package piscine

func TrimAtoi(s string) int {
	r := false
	count := 0
	for _, i := range s {
		if i == '-' && count == 0 {
			r = true
		} else if i >= '0' && i <= '9' {
			count = count*10 + int(i-48)
		}
	}
	if r {
		return -count
	} else {
		return count
	}
}
