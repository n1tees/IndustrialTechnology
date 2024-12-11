package main

import (
    "fmt"
    "strings"
)

func main() {

    for {
        var input string

        var fcouple [2]float32
        var scouple [2]float32
        var thcouple [2]float32

        println("Enter first coupple")
        fmt.Scanln(&fcouple[0], &fcouple[1])
        strings.ToLower(input)

        println("Enter second coupple")
        fmt.Scanln(&scouple[0], &scouple[1])
        strings.ToLower(input)

        println("Enter third coupple")
        fmt.Scanln(&thcouple[0], &thcouple[1])
        strings.ToLower(input)

        fmt.Println(intersection(fcouple[0], fcouple[1], scouple[0], scouple[1], thcouple[0], thcouple[1]))

    }

}

func intersection(a, b, c, d, e, f float32) bool {

    points := []float32{a, b, c, d, e, f}
    for i := 0; i < len(points); i++ {
        for j := i + 1; j < len(points); j++ {
            if points[i] > points[j] {
                points[i], points[j] = points[j], points[i]
            }
        }
    }

    minMax := points[2]
    maxMin := points[3]
    return minMax <= maxMin
}
