package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var s string
	fmt.Print("Введите строку: ")
	fmt.Scanf("%s", &s)
	if isPalindrome(s) {
		fmt.Println("Это палиндром.")
	} else {
		fmt.Println("Это не палиндром.")
	}
}
