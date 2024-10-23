package main

import (
    "fmt"
)

func isArmstrongNumber(n int) bool {
    num := n
    sum := 0
    for num > 0 {
        digit := num % 10
        sum += digit * digit * digit
        num /= 10
    }
    return n == sum
}

func armstrongNumbersInRange(a, b int) []int {
    var armstrongs []int
    for i := a; i <= b; i++ {
        if isArmstrongNumber(i) {
            armstrongs = append(armstrongs, i)
        }
    }
    return armstrongs
}

func main() {
    var a, b int
    fmt.Print("Введите два числа: ")
    fmt.Scanln(&a, &b)

    armstrongs := armstrongNumbersInRange(a, b)
    fmt.Println("Числа Армстронга в диапазоне:")
    for _, armstrong := range armstrongs {
        fmt.Println(armstrong)
    }
}
