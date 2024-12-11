package main

import (
	"fmt"
)
func sumArray(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func main() {
	var n int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Сумма элементов массива:", sumArray(arr))
}
