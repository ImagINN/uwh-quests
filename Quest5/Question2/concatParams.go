package main

import "fmt"

func ConcatParams(args []string) string {
	newStr := ""
	for i, r := range args {
		if i != 0 && i != len(args) {
			newStr = newStr + "\\" + "n"
		}

		newStr += r
	}

	return newStr
}

func main() {
	fmt.Println(ConcatParams([]string{"Hello", "World", "!"}))
}
