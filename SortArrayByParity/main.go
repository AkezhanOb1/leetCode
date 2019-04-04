package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

func sortArrayByParity(A []int) []int {
	var indexCounter int
	var indexVal int
	for i, v := range A {
		if v%2 == 0 {
			indexVal = A[indexCounter]
			A[indexCounter] = v
			A[i] = indexVal
			indexCounter++
		}
	}

	return A
}
