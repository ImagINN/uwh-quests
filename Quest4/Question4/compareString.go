package main

import (
	"fmt"
	"unicode/utf8"
)

func Compare(a, b string) int {
	x := utf8.RuneCount([]byte(a))
	y := utf8.RuneCount([]byte(b))
	if x > y || (a != "" && b == "") {
		return 1
	}
	if x < y || (a == "" && b != "") {
		return -1
	}
	return 0
}

func main() {
	value1 := "Hello!"
	value2 := "Hello"

	compare := Compare(value1, value2)
	fmt.Printf("Compare(%s, %s) => %d\n", value1, value2, compare)
}
