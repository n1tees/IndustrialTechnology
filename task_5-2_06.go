package main

import (
	"fmt"
)
func generatePrimes(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	var limit int
	fmt.Print("Введите предел для генерации простых чисел: ")
	fmt.Scanf("%d", &limit)
	fmt.Println("Простые числа:", generatePrimes(limit))
}
