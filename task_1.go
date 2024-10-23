package main

import (
    "bufio"
    "errors"
    "fmt"
    "math/big"
    "os"
    "strconv"
    "strings"
)

func main() {

    fmt.Println("enter num, currnet system num and target system")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()

    str := strings.Split(input, " ")
    value := str[0]

    var number [2]int
    for i := 1; i < len(str); i++ {
        var err error
        number[i-1], err = strconv.Atoi(str[i])

        if err != nil {
            fmt.Println("Error during conversion, try to enter values again")
            return
        }
    }

    var err error
    output, err := convertBase(value, number[0], number[1])

    if err == nil {
        fmt.Print("your number in target system is " + output)
        return
    } else { fmt.Println(err) }
}
func convertBase(number string, fromBase int, toBase int) (string, error) {
    if fromBase < 2 || fromBase > 36 || toBase < 2 || toBase > 36 {
        return "", errors.New("invalid number system")
    }

    n, ok := new(big.Int).SetString(number, fromBase)
    if !ok {
        return "", errors.New("not possible to convert this number")
    }
    return n.Text(toBase), nil
}
