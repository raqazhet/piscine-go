package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for index, value := range arg {
		if index >= 1 {
			for r := 0; r < len(value); r++ {
				qr := []rune(value)
				z01.PrintRune(qr[r])
			}
			z01.PrintRune('\n')
		}
	}
}
