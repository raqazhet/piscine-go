package piscine

func Map(f func(int) bool, a []int) []bool {
	var raz []bool
	for _, w := range a {
		raz = append(raz, f(w))
	}
	return raz
}
