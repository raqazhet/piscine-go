package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		file, _ := ioutil.ReadAll(os.Stdin)
		for _, i := range string(file) {
			z01.PrintRune(rune(i))
		}
	}
	for _, i := range arg {
		content, err := ioutil.ReadFile(i)
		if err != nil { // ERROR: open abc: no such file or directory
			for _, s := range "ERROR: open " + i + ": no such file or directory\n" {
				z01.PrintRune(s)
			}
			os.Exit(1)
		}
		for _, j := range string(content) {
			z01.PrintRune(j)
		}
	}
}
