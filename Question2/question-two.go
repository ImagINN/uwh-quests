package main

import (
	"fmt"
)

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}

func main() {
	number := 6
	factorial := RecursiveFactorial(number)
	fmt.Printf("%d! = %d\n", number, factorial)
}
