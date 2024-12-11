package main

import (
    "fmt"
    "regexp"
    "strings"
)

func main() {

    for {
        var input string
        println("Enter palindrom")

        fmt.Scanln(&input)
        strings.ToLower(input)

        re := regexp.MustCompile(`[ \n\t,!.?:;]`)
        input = re.ReplaceAllString(input, "")

        for i := 0; i < len(input)/2; i++ {
            if input[i] != input[len(input)-i-1] {
                println("Not palindrom")
                return
            }
        }
        println("Palindrom")

    }

}

