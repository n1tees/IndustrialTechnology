package main

import (
	"fmt"
)
func sumOfSquares(n int) int {
	sum := 0
	for i := 2; i <= n; i += 2 {
		sum += i * i
	}
	return sum
}
func main() {
	var n int
	fmt.Print("Введите число n: ")
	fmt.Scanf("%d", &n)
	fmt.Println("Сумма квадратов чётных чисел:", sumOfSquares(n))
}
