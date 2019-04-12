package main

import "fmt"

func main() {
	A := []int{1, 2, 3, 4}
	queries := [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}
	fmt.Println(sumEvenAfterQueries(A, queries))
}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	var counter int
	var res []int
	for i, v := range queries {
		A[v[1]] = A[v[1]] + queries[i][0]
		counter = evenCounter(A)
		res = append(res, counter)
	}

	return res
}

func evenCounter(A []int) int {
	var evenCounter int
	for _, v := range A {
		if v%2 == 0 {
			evenCounter += v
		}
	}
	return evenCounter
}
