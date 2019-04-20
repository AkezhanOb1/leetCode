package main

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{3, 4, 2, 3}))
}

func isMonotonic(A []int) bool {
	if len(A) == 1 {
		return true
	}
	if A[0] == A[len(A)-1] {
		if A[0] == A[1] {
			return true
		}
		return false
	} else if A[0] > A[len(A)-1] {
		return checkDec(A)
	} else {
		return checkInc(A)
	}
}

func checkInc(A []int) bool {
	checker := true
	for i := 1; i < len(A); i++ {
		if A[i-1] > A[i] {
			checker = false
		}
	}
	return checker
}

func checkDec(A []int) bool {
	checker := true
	for i := 1; i < len(A); i++ {
		if A[i-1] < A[i] {
			checker = false
		}
	}
	return checker
}
