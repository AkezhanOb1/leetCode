package main

import (
	"fmt"
)

func main() {
	fmt.Println(combinationSum3(3, 7))
}

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	for j := 0; j < k; j++ {
		row := []int{}
		lastVal := n - k - j
		rest := sum(n - lastVal, lastVal)
		row = append(row, rest...)
		row = append(row, lastVal)
		if len(row) == k {
			res = append(res, row)			
		}
	}
	return  res
}

func sum(n int, last int) []int {
	counter := map[int]int{}
	res := []int{}
	for i := 0; i < n - 1; i++ {
		if val, ok := counter[i]; ok == true && val != last && i != last {
			res = append(res, val, i)
			return res
		}
		dif := n - i 
		counter[dif] = i
	}
	return res
}