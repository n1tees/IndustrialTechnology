package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Введите количество строк: ")
	fmt.Scan(&n)
	arr := make([]string, n)
	fmt.Println("Введите строки:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var result string
	for i := 0; i < n; i++ {
		result += arr[i]
		if i < n-1 {
			result += " "
		}
	}
	fmt.Println("Объединенная строка:", result)
}
