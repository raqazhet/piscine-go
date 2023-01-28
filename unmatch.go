package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		k := 0
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] && i != j {
				a[i] = -1
				a[j] = -1
				k++
			}
		}
		if k == 0 {
			return a[i]
		}
	}
	return -1
}
