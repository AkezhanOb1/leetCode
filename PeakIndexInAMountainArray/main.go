package main

import "fmt"

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 2, 1, 0}))
}

func peakIndexInMountainArray(A []int) int {
	max := A[0]
	index := 0

	for i, v := range A {
		if v > max {
			max = v
			index = i
		}
	}
	return index
}
