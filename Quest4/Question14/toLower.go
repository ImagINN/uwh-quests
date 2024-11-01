package main

import (
	"fmt"
	"unicode"
)

func ToLower(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if unicode.IsUpper(r) {
			runes[i] = unicode.ToLower(r)
		}
	}
	s = string(runes)

	return s
}

func main() {
	fmt.Println(ToLower("HELLO, World!"))
}
