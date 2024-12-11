package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2, num3 int
	fmt.Print("Введите три числа: ")
	fmt.Scan(&num1, &num2, &num3)
	max := num1
	if num2 > max {
		max = num2
	}
	if num3 > max {
		max = num3
	}
	fmt.Printf("Наибольшее число: %d\n", max)
}
