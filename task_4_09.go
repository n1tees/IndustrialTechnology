package main

import (
	"fmt"
	"math"
)


func main() {
	var age int
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)
	if age < 14 {
		fmt.Println("ребенок")
	} else if age < 18 {
		fmt.Println("подросток")
	} else if age < 65 {
		fmt.Println("взрослый")
	} else {
		fmt.Println("пожилой")
	}
}
