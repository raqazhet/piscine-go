package piscine

func ActiveBits(n int) int {
	var slice []int
	count := 0
	for i := 0; i < 8; i++ {
		slice = append(slice, n%2)
		if n%2 == 1 {
			count++
		}
		n /= 2
	}
	return count
}
