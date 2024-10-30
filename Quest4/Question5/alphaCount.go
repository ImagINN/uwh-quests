/*
Instructions
Write a function that counts the letters of a string and returns the count.

The letters are from the Latin alphabet list only, any other characters, symbols or empty spaces shall not be counted.
*/

package main

import (
	"fmt"
)

func AlphaCount(s string) int {
	count := 0
	for _, r := range s {
		if r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' {
			count++
		}
	}
	return count
}

func main() {
	value := "Hello, world!"
	count := AlphaCount(value)
	fmt.Printf("Number of letters in %s is %d\n", value, count)
}
