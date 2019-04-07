package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayPairSum([]int{1, 1}))
}

func arrayPairSum(nums []int) int {
	sort.IntSlice(nums).Sort()
	counter := 0

	for i := 0; i < len(nums); i = i + len(nums)/(len(nums)/2) {
		counter = counter + nums[i]
	}

	return counter
}
