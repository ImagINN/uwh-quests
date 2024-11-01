package main

import (
	"fmt"
	"unicode"
)

func IsPrintable(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if unicode.IsControl(r) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(IsPrintable("123"))
	fmt.Println(IsPrintable("123abc"))
	fmt.Println(IsPrintable("123 123"))
	fmt.Println(IsPrintable("\x00"))
}
