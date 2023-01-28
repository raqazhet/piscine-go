package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	for i := 0; i <= 9; i++ {
		a := n
		if a == 0 {
			z01.PrintRune(48)
			break
		}
		for ; a > 0; a /= 10 {
			ost := a % 10
			if i == ost {
				z01.PrintRune(rune(i + 48))
			}

		}
	}
}
