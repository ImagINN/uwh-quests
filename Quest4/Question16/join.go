package main

import (
	"fmt"
)

func Join(strs []string, sep string) string {
	newStr := ""
	for i, r := range strs {
		if i != 0 && i != len(strs) {
			newStr += sep
		}

		newStr += r
	}

	return newStr
}

func main() {
	strs := []string{"Hello", "World", "Golang"}
	fmt.Println(Join(strs, ", "))
}
