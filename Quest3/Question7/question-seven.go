package main

import (
	"fmt"
)

func IsPrime(nb int) bool {
	if nb <= 0 || nb == 1 {
		return false
	}

	if nb == 2 {
		return true
	} else if nb%2 == 0 {
		return false
	}

	for i := 3; i <= nb/2; i += 2 {
		if nb%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	number := 5
	value := IsPrime(number)
	fmt.Printf("Is %d a prime number? %t\n", number, value)
}
