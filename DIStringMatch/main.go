package main

import "fmt"

func main() {
	fmt.Println(diStringMatch("DDI"))
}

func diStringMatch(S string) []int {
	max := len(S)
	min := 0
	var res []int

	for _, v := range S {
		if v == 'I' {
			res = append(res, min)
			min++
		} else {
			res = append(res, max)
			max--
		}
	}

	res = append(res, min)

	return res
}
