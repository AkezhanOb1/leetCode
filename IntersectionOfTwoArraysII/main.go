package main

import "fmt"

func main() {
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	mapper := make(map[int]int8)
	for _, num := range nums1 {
		if _, ok := mapper[num]; !ok {
			mapper[num]++
		} else {
			mapper[num]++
		}
	}
	for _, num := range nums2 {
		if val, ok := mapper[num]; ok && val > 0 {
			mapper[num]--
			res = append(res, num)
		}
	}

	return res
}
