package main

import "fmt"

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    var n int
    fmt.Print("Введите целое число: ")
    fmt.Scanln(&n)

    fmt.Println("Последовательность чисел Фибоначчи:")
    for i := 0; ; i++ {
        fib := fibonacci(i)
        if fib > n {
            break
        }
        fmt.Println(fib)
    }
}

