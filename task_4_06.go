package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Printf("%d является четным.\n", number)
	} else {
		fmt.Printf("%d является нечетным.\n", number)
	}
}
