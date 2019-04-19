package main

import "fmt"

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	mapper := make(map[int]bool)
	for _, num := range nums1 {
		if _, ok := mapper[num]; !ok {
			mapper[num] = false
		}
	}
	for _, num := range nums2 {
		if val, ok := mapper[num]; ok && val == false {
			mapper[num] = true
			res = append(res, num)
		}
	}

	return res
}
