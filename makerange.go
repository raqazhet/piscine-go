package piscine

func MakeRange(min, max int) []int {
	var size int = max - min
	if size < 1 {
		return []int(nil)
	}
	make := make([]int, size)
	for i := min; i < max; i++ {
		make[i-min] = i
	}
	return make
}
