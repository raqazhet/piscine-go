package piscine

func Compact(ptr *[]string) int {
	var raz []string
	count := 0
	for _, v := range *ptr {
		if v != "" {
			raz = append(raz, v)
			count++
		}
		*ptr = raz
	}
	return count
}
