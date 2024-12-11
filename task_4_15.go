package main

import (
	"fmt"
	"math"
)

func main() {
	var numbers int
	fmt.Print("Введите количество элементов массива: ")
	fmt.Scan(&numbers)
	arr := make([]int, numbers)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < numbers; i++ {
		fmt.Scan(&arr[i])
	}
	sum := 0
	for _, number := range arr {
		sum += number
