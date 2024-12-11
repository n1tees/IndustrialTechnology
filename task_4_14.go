package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	primes := []int{}
	for i := 2; i <= n; i++ {
		prime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			primes = append(primes, i)
		}
	}
	fmt.Println("Простые числа до", n, ":", primes)
}
