package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDuplicate([]int{1,3,4,2,2}))
}

func findDuplicate(nums []int) int {
	counter := map[int]int{}
	for _, num := range nums {
		if _, ok := counter[num]; ok {
			return num
		}
		counter[num] = 1
		
	}
	return 0
}