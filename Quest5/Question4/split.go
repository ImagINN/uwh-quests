package main

import (
	"fmt"
	"strings"
)

func Split(s, sep string) []string {
	var result []string
	start := 0

	for {
		index := strings.Index(s[start:], sep)
		if index == -1 {
			result = append(result, s[start:])
			break
		}

		result = append(result, s[start:start+index])
		start += index + len(sep)
	}

	return result
}

func main() {
	fmt.Println(Split("a,b,c", ","))
	fmt.Println(Split("hello world hello", " "))
	fmt.Println(Split("one::two::three", "::"))
	fmt.Println(Split("nospace", ","))
}
