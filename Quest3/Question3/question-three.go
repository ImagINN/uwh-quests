package main

import (
	"fmt"
)

func IterativePower(nb, power int) int {
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	if power == 1 {
		return nb
	}

	res := 1
	for i := 1; i <= power; i++ {
		res *= nb
	}

	return res
}

func main() {
	number := 2
	power := 3
	factorial := IterativePower(number, power)
	fmt.Printf("%d! = %d\n", number, factorial)
}
