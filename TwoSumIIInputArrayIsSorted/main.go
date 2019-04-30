package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
}

func twoSum(numbers []int, target int) []int {
	prevIndex := 0
	res := []int{}
	for prevIndex < len(numbers) {
		for i := prevIndex + 1; i < len(numbers); i++ {
			if numbers[prevIndex]+numbers[i] == target {
				res = append(res, []int{prevIndex + 1, i + 1}...)
				return res
			}
		}
		prevIndex++
	}
	return res
}
