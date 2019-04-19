package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}

func singleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result = result ^ v
	}
	return result
}
