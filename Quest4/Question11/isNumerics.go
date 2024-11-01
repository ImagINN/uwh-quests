package main

import (
	"fmt"
	"unicode"
)

func IsNumeric(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(IsNumeric("123"))
	fmt.Println(IsNumeric("123abc"))
	fmt.Println(IsNumeric("123 123"))
	fmt.Println(IsNumeric("123, 123"))
}
