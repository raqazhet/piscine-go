package main

import (
	"fmt"
)

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

func Atoi(s string) int {
	var num int
	raz := []rune(s)
	negative := 1
	if raz[0] == '-' {
		negative *= -1
		raz = raz[1:]
	} else if raz[0] == '+' {
		raz = raz[1:]
	}
	for _, v := range raz {
		if v < '0' || v > '9' {
			num = 0
			break
		}

		num = num*10 + int(v-48)
	}
	return num * negative
}