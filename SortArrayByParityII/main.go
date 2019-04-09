package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
}

func sortArrayByParityII(A []int) []int {
	wrongOdd := []int{}
	wrongEven := []int{}
	for i, v := range A {
		if i%2 == 0 && v%2 != 0 {
			wrongOdd = append(wrongOdd, i)
		}
		if i%2 != 0 && v%2 == 0 {
			wrongEven = append(wrongEven, i)
		}
	}

	for i := range wrongEven {
		midVal := A[wrongEven[i]]
		A[wrongEven[i]] = A[wrongOdd[i]]
		A[wrongOdd[i]] = midVal
	}
	return A
}
