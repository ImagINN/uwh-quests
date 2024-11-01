package main

import (
	"fmt"
	"unicode"
)

func ToUpper(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if unicode.IsLower(r) {
			runes[i] = unicode.ToUpper(r)
		}
	}
	s = string(runes)

	return s
}

func main() {
	fmt.Println(ToUpper("Hello, World!"))
}
