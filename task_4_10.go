package main

import (
	"fmt"
	"math"
)


func main() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	if number%3 == 0 && number%5 == 0 {
		fmt.Printf("Число %d делится на 3 и 5.\n", number)
	} else {
		fmt.Printf("Число %d не делится одновременно на 3 и 5.\n", number)
	}
}
