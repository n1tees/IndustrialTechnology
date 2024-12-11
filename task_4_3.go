package main

import (
	"fmt"
	"math"
)

func main() {
	var size int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&size)

	numbers := make([]int, size)
	fmt.Println("Введите числа:")
	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
	}

	doubled := make([]int, len(numbers))
	for i, num := range numbers {
		doubled[i] = num * 2
	}

	fmt.Println("Удвоенные числа:", doubled)
}
