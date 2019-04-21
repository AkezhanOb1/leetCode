package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(rotatedDigits(66))
}

func rotatedDigits(N int) int {
	counter := 0
	for i := 1; i <= N; i++ {
		if ok := rotatableNum(i); ok {
			counter++
		}
	}
	return counter
}

func rotatableNum(N int) bool {
	strN := strconv.Itoa(N)
	rotMap := map[rune]rune{
		'0': '0',
		'1': '1',
		'2': '5',
		'5': '2',
		'6': '9',
		'8': '8',
		'9': '6',
	}
	resStr := ""
	for _, num := range strN {
		if _, ok := rotMap[num]; !ok {
			return false
		}
		resStr += string(rotMap[num])
	}
	resInt, _ := strconv.Atoi(resStr)
	if resInt == N {
		return false
	}
	return true
}

func isValid(n int) bool {

	isChanged := false
	for n > 0 {

		switch n % 10 {
		case 3, 4, 7:
			return false
		case 2, 5, 6, 9:
			isChanged = true
		}
		n /= 10
	}

	return isChanged
}
