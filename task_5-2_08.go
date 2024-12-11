package main

import (
	"fmt"
)
func findMax(arr []int) int {
	maxT := arr[0]
	for _, value := range arr {
		if value > maxT {
			maxT = value
		}
	}
	return maxT
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
	fmt.Println("Максимальный элемент:", findMax(arr))
}

