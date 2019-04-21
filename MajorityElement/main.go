package main

import (
	"fmt"
)

func main() {
	fmt.Printf("The majority element is %d\n", majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElement(nums []int) int {
	majority := len(nums) / 2
	majorityElement := 0
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
			majorityElement = key
		}
	}

	return majorityElement
}

// SECOND APPROACH
// func majorityElement(nums []int) int {
// 	sort.Ints(nums)
// 	counter := 1
// 	majorityElement := nums[len(nums)-1]
// 	indexer := 0
// 	for i := 1; i < len(nums); i++ {
// 		if nums[indexer] != nums[i] {
// 			if counter > len(nums)/2 {
// 				majorityElement = nums[indexer-1]
// 			}
// 			indexer = i
// 			counter = 0
// 		}
// 		counter++
// 	}
// 	return majorityElement
// }

// THIRD APPROACH
// func majorityElement(nums []int) int {
// 	sort.Ints(nums)
// 	return nums[len(nums) / 2]
// }

// FOURTH APPROACH
// func majorityElement(nums []int) int {
// 	sm := make(map[int]int)
// 	for _, n := range nums {
// 		m[n]++
// 		if m[n] > len(nums)/2 {
// 			return n
// 		}
// 	}
// 	return 0
// }
