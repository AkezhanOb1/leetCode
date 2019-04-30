package main

import (
	"fmt"
)

func main() {
	fmt.Printf("The majority element is %d\n", majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElement(nums []int) []int {
	majority := len(nums) / 3
	res := []int{}
	counterMap := make(map[int]int)
	for _, num := range nums {
		if _, ok := counterMap[num]; ok {
			counterMap[num]++
		} else {
			counterMap[num] = 1
		}
	}

	for key, val := range counterMap {
		if val > majority {
			res = append(res, key)
		}
	}

	return res
}
