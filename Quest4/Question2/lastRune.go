/*
Instructions:
Write a function that returns the first rune of a string.
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func LastRune(s string) rune {
	r, _ := utf8.DecodeRuneInString(s[len(s)-1:])
	return r
}

func main() {
	value := "Hello"
	firstRune := LastRune(value)
	fmt.Printf("Last rune of %s is %c\n", value, firstRune)
}
