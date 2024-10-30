package main

import (
	"fmt"
)

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}

func main() {
	index := 6
	value := Fibonacci(index)
	fmt.Printf("Fibonacci value at index %d is %d\n", index, value)
}
