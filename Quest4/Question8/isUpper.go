package main

import "fmt"

func IsUpper(s string) bool {
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsUpper("Hello, World!"))
	fmt.Println(IsUpper("HELLO, WORLD!"))
	fmt.Println(IsUpper("HELLO"))
}
