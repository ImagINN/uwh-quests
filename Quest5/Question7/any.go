package main

import "fmt"

func Any(f func(string) bool, a []string) bool {
	if len(a) == 0 {
		return false
	}

	for _, n := range a {
		if f(n) {
			return true
		}
	}
	return false
}

func IsNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func main() {

	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
