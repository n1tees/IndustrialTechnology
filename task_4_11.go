package main

import (
	"fmt"
	"math"
)


func task11() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if number < 0 {
		fmt.Println("Нет факториал отрицательного числа")
	} else {
		result := 1
		for i := 1; i <= number; i++ {
			result *= i
		}
		fmt.Printf("Факториал числа %d равен %d\n", number, result)
	}
}
