package main

import (
	"fmt"
	"math"
)


func main() {
	var number int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&number)
	if number <= 0 {
		fmt.Println("Введено отрицательное число")
	} else {
		fib := make([]int, number)
		if number > 0 {
			fib[0] = 0
		}
		if number > 1 {
			fib[1] = 1
		}
		for i := 2; i < number; i++ {
			fib[i] = fib[i-1] + fib[i-2]
		}
		fmt.Printf("Первые %d чисел Фибоначчи: %v\n", number, fib)
	}
}
