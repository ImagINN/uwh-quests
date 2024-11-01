package main

import (
	"fmt"
	"strings"
)

func SplitWhiteSpaces(s string) []string {
	return strings.Fields(s)
}

func main() {
	fmt.Println(SplitWhiteSpaces("Hello World from Go!"))
	fmt.Println(SplitWhiteSpaces("   Multiple   spaces    here  "))
	fmt.Println(SplitWhiteSpaces(""))
}

