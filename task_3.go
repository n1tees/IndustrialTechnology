package main

import (
    "fmt"
    "math"
    "math/rand"
    "strconv"
    "time"
)

func main() {

    fmt.Println("enter the count of num in mass")
    var input string
    fmt.Scanln(&input)

    var value, err = strconv.Atoi(input)

    if err != nil {
        fmt.Println("error")
        return
    }

    var mass []int = make([]int, value)
    fillMass(mass)
    sortMass(mass)

    fmt.Println(mass)
}

func fillMass(mass []int) {
    time.Now().UnixNano()
    for i := range mass {
        mass[i] = rand.Intn(201) - 100
    }
    fmt.Print(mass, "\n")
}

func sortMass(mass []int) {
    for i := 0; i < len(mass); i++ {
        mass[i] = int(math.Abs(float64(mass[i])))
    }

    for i := 0; i < len(mass); i++ {
        for j := i + 1; j < len(mass); j++ {
            if mass[i] > mass[j] {
                mass[i], mass[j] = mass[j], mass[i]
            }
        }
    }

}
