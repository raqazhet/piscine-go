package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	raz := os.Args[1:]
	if len(raz) != 3 {
		return
	}
	a := Atoi(raz[0])
	// if err != nil {
	// 	return
	// }
	b := Atoi(raz[2])
	// if err != nil {
	// 	return
	// }

	// OSYMEN ZHYMYS ISTEIDI
	maxint := int(^uint(0) >> 1)
	minint := -maxint - 1
	if (a > 0 && b > 0 && a+b < 0) || a > maxint || b > maxint || a < minint || b < minint {
		return
	}
	// jsvkglcihajs;dvupoiurhwipu yn8pwrygpwy guwy guwhgwhou hgwo ygo8yhgouwh ohwoghw oihwoih
	// RAZAQ
	// Ystindegi kotty qara

	// if a > 0 && b > 0 && (a+b) < 0 || (a < 0 && b < 0 && (a+b) > 0) {
	// 	return
	// }
	switch raz[1] {
	case "+":
		PrintStr(Itoa(a + b))
	case "-":
		PrintStr(Itoa(a - b))
	case "*":
		PrintStr(Itoa(a * b))
	case "/":
		if b == 0 {
			PrintStr("No division by 0")
			return
		}
		PrintStr(Itoa(a / b))
	case "%":
		if b == 0 {
			PrintStr("No modulo by 0")
			return
		}
		PrintStr(Itoa(a % b))
	}
}

func Atoi(str string) int {
	num := 0
	negative := 1
	raz := []rune(str)
	if raz[0] == '-' {
		negative = -1
		raz = raz[1:]
	}
	for _, v := range raz {
		if v < '0' || v > '9' {
			os.Exit(0)
		}
		num = num*10 + int(v-48)
	}
	return num * negative
}

func Itoa(n int) string {
	str := ""
	for n >= 10 || n <= -10 {
		d := n % 10
		if d < 0 {
			d *= -1
		}
		str = string(rune(d+48)) + str
		n /= 10
	}
	if n > 0 {
		str = string(rune(n+48)) + str
	} else {
		n *= -1
		str = "-" + string(rune(n+48)) + str
	}
	return str
}

func PrintStr(str string) {
	raz := []rune(str)
	for i := 0; i < len(raz); i++ {
		z01.PrintRune(raz[i])
	}
	z01.PrintRune('\n')
}
