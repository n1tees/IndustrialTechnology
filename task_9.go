package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)

func main() {
    var input string
    scanner := bufio.NewReader(os.Stdin)
    input, _ = scanner.ReadString('\n')

    wordsSlice := strings.Split(input, " ")

    re := regexp.MustCompile(`[\n\t,!.?:;]`)

    var maxWord string
    var maxLen = -1

    for i := 0; i < len(wordsSlice); i++ {
        word := re.ReplaceAllString(wordsSlice[i], "")
        if len(word) > maxLen {
            maxWord = word
            maxLen = len(word)
        }
    }

    fmt.Print(maxWord)
    fmt.Scanln()
}
