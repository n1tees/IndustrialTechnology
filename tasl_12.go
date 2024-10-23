package main

import "fmt"

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func primeNumbersInRange(a, b int) []int {
    var primes []int
    for i := a; i <= b; i++ {
        if isPrime(i) {
            primes = append(primes, i)
        }
    }
    return primes
}

func main() {
    var a, b int
    fmt.Print("Введите два числа: ")
    fmt.Scanln(&a, &b)

    primes := primeNumbersInRange(a, b)
    fmt.Println("Простые числа в диапазоне:")
    for _, prime := range primes {
        fmt.Println(prime)
    }
}

