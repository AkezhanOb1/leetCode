package main

import (
	"fmt"
)

func main() {
	fmt.Println(rotateString("abcde", "cdeab"))
}
func rotateString(A string, B string) bool {
	for i := 0; i < len(A); i++ {
		fmt.Println(A[i:]+A[:i])
		if A[i:]+A[:i] == B {
			return true
		}
	}
	return A == B
}