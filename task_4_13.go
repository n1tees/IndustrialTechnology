package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	fmt.Print("Введите количество элементов массива: ")
	fmt.Scan(&number)
	arr := make([]int, number)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < number; i++ {
		fmt.Scan(&arr[i])
	}
	reversed := make([]int, number)
	for i := 0; i < number; i++ {
		reversed[i] = arr[number-i-1]
	}
	fmt.Println("Реверсированный массив:", reversed)
}
