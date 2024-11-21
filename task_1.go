	package main
	
	import (
	    "fmt"
	    "strconv"
	)
	func main() {
	
	    fmt.Print("enter temperature for convert\n")
	    var s1, s2 string
	    var value int
	    fmt.Scan(&s1, &s2)
	
	    value, err := strconv.Atoi(s1)
	
	    if err != nil {
	        fmt.Println("Error during conversion")
	        return
	    }
	
	    switch s2 {
	
	    case "(Celsius)":
	        fmt.Print((float32(value) * 9 / 5) + 32)
	        fmt.Print(" Fahrenhiet")
	    case "(Fahrenheit)":
	        fmt.Print((float32(value) - 32) * 5 / 9)
	        fmt.Print(" Celsius")
	    default:
	        fmt.Print("undefined value")
	    }
	}
