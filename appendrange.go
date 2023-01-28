package piscine

func AppendRange(min, max int) []int {
	var massiv []int
	if min >= max {
		return nil
	} else {
		for i := min; i < max; i++ {
			massiv = append(massiv, i)
		}
		return massiv
	}
}
