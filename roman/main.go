package main

import (
	"fmt"
	"os"
	"strconv"
)

type roman struct {
	l string
	n int
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]
	num, err := strconv.Atoi(args)
	if err != nil {
	}
	array := []roman{{"I", 1}, {"IV", 4}, {"V", 5}, {"IX", 9}, {"X", 10}, {"XL", 40}, {"L", 50}, {"XC", 90}, {"C", 100}, {"CD", 400}, {"D", 500}, {"CM", 900}, {"M", 1000}}
	//					0			1		2			3			4			5			6 		7			8				9		  10		   11  	 		12
	res := ""

	for i := len(array) - 1; i >= 0; i-- {
		if num/array[i].n != 0 {
			c := num / array[i].n
			for j := 0; j < c; j++ {
				res += array[i].l
			}
			num %= array[i].n
		}
	}
	fmt.Println(res)
}
