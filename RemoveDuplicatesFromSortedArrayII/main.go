package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 2}))
}

func removeDuplicates(nums []int) int {
	current, previous := 2, 1

	for current < len(nums) {
		if nums[current] == nums[previous] && nums[previous] == nums[previous-1] {
			current++
		} else {
			previous++
			nums[previous] = nums[current]
			current++
		}
	}
	return previous + 1
}
