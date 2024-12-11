package main

import (
	"fmt"
)
unc isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scanf("%d", &n)
	if isPrime(n) {
		fmt.Println("Число простое.")
	} else {
		fmt.Println("Число не является простым.")
	}
}

