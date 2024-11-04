package main

import "fmt"

func ForEach(f func(int), a []int) {
	for _, n := range a {
		f(n)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := ""
	ForEach(func(n int) {
		result += fmt.Sprintf("%d", n)
	}, arr)
	fmt.Println(result)
}
