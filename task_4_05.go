package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Print("Введите координаты первой точки (x1, y1): ")
	fmt.Scan(&x1, &y1)
	fmt.Print("Введите координаты второй точки (x2, y2): ")
	fmt.Scan(&x2, &y2)
	distance := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	fmt.Printf("Расстояние между точками (%.1f, %.1f) и (%.1f, %.1f) равно %.1f\n", x1, y1, x2, y2, distance)
}
