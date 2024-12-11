package main

import (
	"fmt"
)
func toBinary(n int) string {
	binary := ""
	for n > 0 {
		binary = fmt.Sprintf("%d", n%2) + binary
		n /= 2
	}
	return binary
}

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scanf("%d", &n)
	fmt.Println("Двоичное представление числа:", toBinary(n))
}
