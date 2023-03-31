package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min // TODO: replace this
}

func FindMax(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func SumMinMax(nums ...int) int {
	min := FindMin(nums...)
	max := FindMax(nums...)

	return min + max
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax( 3, 1, 4, 5, 6, 7, 8, 9, 13))
}
