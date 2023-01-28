package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, w := range tab {
		if f(w) {
			count++
		}
	}
	return count
}
