package piscine

func BasicJoin(elems []string) string {
	res := ""
	for i := range elems {
		res += elems[i]
	}
	return res
}
