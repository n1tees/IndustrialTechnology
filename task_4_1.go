package main

import (
	"fmt"
	"math"
)
func main() {
	var number int
	fmt.Print("Введите целое число: ")
	fmt.Scan(&number)
	if number < 0 {
		number = -number
	}
	sum := sumDigits(number)
	fmt.Println("Сумма цифр числа:", sum)
}

func sumDigits(number int) int {
	sum := 0
	for number > 0 {
		sum += number % 10
		number /= 10
	}
	return sum
}
