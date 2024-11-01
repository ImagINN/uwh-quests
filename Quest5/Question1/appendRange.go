package main

import "fmt"

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}

	nums := make([]int, 0, max-min)
	for i := min; i < max; i++ {
		nums = append(nums, i)
	}

	return nums
}

func main() {
	fmt.Println(AppendRange(1, 5))
	fmt.Println(AppendRange(1, 1))
	fmt.Println(AppendRange(5, 1))
	fmt.Println(AppendRange(3, 4))
	fmt.Println(AppendRange(0, 0))
}
