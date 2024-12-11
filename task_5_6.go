package main

import (
    "fmt"
    "math"
    "strconv"
)

func main() {

    for {
        var input string
        println("Enter first number")

        fmt.Scanln(&input)
        fnum, err := strconv.ParseFloat(input, 10)

        println("Enter second number")

        fmt.Scanln(&input)
        snum, err := strconv.ParseFloat(input, 10)

        if err != nil {
            fmt.Println(err)

        }
        println("+, -, *, /, ^, %")
        fmt.Scanln(&input)

        switch input {
        case "+":
            fmt.Println(fnum + snum)
        case "-":
            fmt.Println(fnum - snum)
        case "*":
            fmt.Println(fnum * snum)
        case "/":
            fmt.Println(fnum / snum)
        case "^":
            fmt.Println(math.Exp(math.Log(fnum) * snum))
        case "%":
            fmt.Println(fnum - math.Floor(fnum/snum)*snum)
        }
    }
}

