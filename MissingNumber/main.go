package main

import (
	"fmt"
)

func main() {
	fmt.Println(missingNumber([]int{0}))
}

func missingNumber(nums []int) int {
	expectedSum := len(nums) * (len(nums) + 1) / 2
	actualSum := 0
	for _, val := range nums {
		actualSum += val
	}

	return expectedSum - actualSum
}
