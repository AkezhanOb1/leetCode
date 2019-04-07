package main

import (
	"fmt"
)

func main() {
	fmt.Println(selfDividingNumbers(220, 255))
}

func selfDividingNumbers(left int, right int) []int {
	var res []int
	for num := left; num <= right; num++ {
		if isDivisible(num) {
			res = append(res, num)
		}
	}
	return res
}

func isDivisible(val int) bool {
	num := val
	for num > 0 {
		digit := num % 10
		if digit == 0 || val%digit != 0 {
			return false
		}
		num = num / 10
	}
	return true
}
