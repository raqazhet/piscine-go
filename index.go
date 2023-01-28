package piscine

func Index(s string, toFind string) int {
	x := len(s)
	r := len(toFind)

	for i := 0; i < x-r; i++ {
		if toFind == s[i:i+r] {
			return i
		}
	}
	return -1
}
