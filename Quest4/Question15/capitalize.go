package main

import (
	"fmt"
	"unicode"
)

func Capitalize(s string) string {
	runes := []rune(s)
	capitalizeNext := true

	for i, r := range runes {
		if capitalizeNext && unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			capitalizeNext = false
		} else {
			runes[i] = unicode.ToLower(r)
		}

		if unicode.IsNumber(r) {
			capitalizeNext = false
		} else if unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) {
			capitalizeNext = true
		}
	}

	return string(runes)
}

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are-things+4you*maTe?"))
}
