package main

import (
	"fmt"
)

func main() {
  	var base, height float64
	fmt.Print("Введите основание треугольника: ")
	fmt.Scanf("%f", &base)
	fmt.Print("Введите высоту треугольника: ")
	fmt.Scanf("%f", &height)
	fmt.Println("Площадь треугольника:", triangleArea(base, height))
}

func triangleArea(base float64, height float64) float64 {
	return 0.5 * base * height
}




  
