package main

import "fmt"

func IsLower(s string) bool {
	if s == "" {
		return false
	} else {
		for _, r := range s {
			if r < 'a' || r > 'z' {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(IsLower("Hello, World!"))
	fmt.Println(IsLower("HELLO, WORLD!"))
	fmt.Println(IsLower("hello"))
}
