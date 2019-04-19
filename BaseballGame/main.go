package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
}

func calPoints(ops []string) int {
	stack := make([]int, len(ops))
	index := 0
	counter := 0

	for _, op := range ops {
		switch op {
		case "+":
			stack[index] = stack[index-1] + stack[index-2]
			counter += stack[index]
			index++
		case "D":
			stack[index] = stack[index-1] * 2
			counter += stack[index]
			index++
		case "C":
			counter -= stack[index-1]
			index--
		default:
			stack[index], _ = strconv.Atoi(op)
			counter += stack[index]
			index++
		}
	}

	return counter
}
