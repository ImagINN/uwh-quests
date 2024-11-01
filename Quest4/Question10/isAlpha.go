package main

import (
	"fmt"
	"unicode"
)

func IsAlpha(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsAlpha("Hello, World!"))
	fmt.Println(IsAlpha("HELLO, WORLD!"))
	fmt.Println(IsAlpha("hello4"))
}
