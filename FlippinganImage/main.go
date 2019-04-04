package main

import "fmt"

func main() {
	a := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	fmt.Println(flipAndInvertImage(a))
}

func flipAndInvertImage(A [][]int) [][]int {
	for outI := range A {
		A[outI] = reverse(A[outI])
		for innI := range A[outI] {
			if A[outI][innI] == 0 {
				A[outI][innI] = 1
			} else {
				A[outI][innI] = 0
			}
		}
	}
	return A
}

func reverse(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
