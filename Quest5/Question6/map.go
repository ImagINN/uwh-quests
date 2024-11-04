package main

import "fmt"

// Instructions
// Write a function Map that, for an int slice, applies a function of this type func(int) bool on each element of that slice and returns a slice of all the return values.

func Map(f func(int) bool, a []int) []bool {
	if len(a) == 0 {
		return []bool{}
	}

	var result []bool
	for _, n := range a {
		result = append(result, f(n))
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := Map(func(n int) bool {
		return n%2 == 0
	}, arr)
	fmt.Println(result)
}
