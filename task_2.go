package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

func main() {

    fmt.Println("enter coefficients of the quadratic equation ax^2 + bx + c = 0")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()

    str := strings.Split(input, " ")

    var number [3]float64
    for i := 0; i < len(str); i++ {
        var err error
        number[i], err = strconv.ParseFloat(str[i], 64)

        if err != nil {
            fmt.Println("Error during conversion, try to enter values again")
            return
        }
    }

    var d float64 = (number[1])*(number[1]) - 4*(number[0])*(number[2])
    fmt.Println(d)
    if d > 0 {
        Root(number, d)
    }
    if d < 0 {
        complexRoot(number, d)
    } else if d == 0 {
        fmt.Println("there is no roots, D=0")
    }
}

func Root(koef [3]float64, d float64) {
    var x1 float64 = (-koef[1] + math.Sqrt(d)) / (2 * koef[0])
    var x2 float64 = (-koef[1] - math.Sqrt(d)) / (2 * koef[0])

    fmt.Println("x1 = ", x1, "\nx2 = ", x2)
}

func complexRoot(koef [3]float64, d float64) {
    var x1 complex128 = (complex(-koef[1], 0) + complex(0, math.Sqrt(-d))) / complex(2*koef[0], 0)
    var x2 complex128 = (complex(-koef[1], 0) - complex(0, math.Sqrt(-d))) / complex(2*koef[0], 0)

    fmt.Println("x1 = ", x1, "\nx2 = ", x2)
}
