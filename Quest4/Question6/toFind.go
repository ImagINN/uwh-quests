package main

import "fmt"

func Index(s, toFind string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			for j := 0; j < len(toFind); j++ {
				if s[i+j] != toFind[j] {
					break
				}
				if j == len(toFind)-1 {
					return i
				}
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(Index("Hello, World!", "World"))
}
