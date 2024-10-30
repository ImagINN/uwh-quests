/*
Instructions:
Write a function that returns the first rune of a string.
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func FirstRune(s string) rune {
	r, _ := utf8.DecodeRuneInString(s)
	return r
}

func main() {
	value := "Hello!"
	firstRune := FirstRune(value)
	fmt.Printf("First rune of %s is %c\n", value, firstRune)
}
