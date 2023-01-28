package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0]
	for index, raz := range arg {
		if index > 1 {
			z01.PrintRune(raz)
		}
	}
	z01.PrintRune('\n')
}
