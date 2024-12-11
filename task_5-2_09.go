package main

import (
	"fmt"
)
func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scanf("%d", &a)
	fmt.Print("Введите второе число: ")
	fmt.Scanf("%d", &b)
	fmt.Println("НОД:", gcd(a, b))
}
