package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(-101))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return reverse(strconv.Itoa(x)) == strconv.Itoa(x)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
