package main

import "fmt"

func main() {
	fmt.Println(repeatedNTimes([]int{2, 1, 2, 5, 3, 2}))
}

func repeatedNTimes(A []int) int {
	var correctNum int
	counter := make(map[int]int)
	for _, row := range A {
		counter[row]++
	}

	for k, v := range counter {
		if v == len(counter)-1 {
			correctNum = k
		}
	}
	return correctNum
}
