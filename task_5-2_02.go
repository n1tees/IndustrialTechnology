package main

import (
	"fmt"
)

func sortArray(arr []float64) []float64 {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	var n int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scanf("%d", &n)
	arr := make([]float64, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%f", &arr[i])
	}
	fmt.Println("Отсортированный массив:", sortArray(arr))
}


  
