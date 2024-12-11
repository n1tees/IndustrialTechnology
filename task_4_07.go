package main

import (
	"fmt"
	"math"
)

func main() {
	var year int
	var isLeapYear bool
	fmt.Print("Введите год: ")
	fmt.Scan(&year)
	isLeapYear = year%4 == 0 && (year%100 != 0 || year%400 == 0)
	if isLeapYear {
		fmt.Printf("%d является високосным годом.\n", year)
	} else {
		fmt.Printf("%d не является високосным годом.\n", year)
	}
}
