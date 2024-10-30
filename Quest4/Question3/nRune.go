/*
Write a function that returns the nth rune of a string. If not possible, it returns 0.
*/

package main

import (
	"fmt"
)

func nRune(s string, n int) rune {
	runes := []rune(s)
	if n <= 0 || n > len(runes) {
		return 0
	}
	return runes[n-1]
}

func main() {
	value := "Hello!"
	nthRune := nRune(value, 2)
	fmt.Printf("nth rune of '%s' is '%c'\n", value, nthRune)

	value2 := "😊👍🎉"
	nthRune2 := nRune(value2, 3)
	fmt.Printf("nth rune of '%s' is '%c'\n", value2, nthRune2)

	nthRuneInvalid := nRune(value, 10)
	fmt.Printf("nth rune of '%s' is '%d'\n", value, nthRuneInvalid)
}
