package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("III"))
}
func romanToInt(s string) int {
	var result, current, previous int
	for i := 0; i < len(s); i++ {
		current = romanValue(s[i])
		result += current
		if previous < current {
			result -= 2 * previous
		}
		previous = current
	}
	return result
}

func romanValue(c byte) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
