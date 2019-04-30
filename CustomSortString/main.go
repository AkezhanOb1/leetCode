package main

import (
	"fmt"
)

func main() {
	fmt.Println(customSortString("cba", "abcd"))
}

func customSortString(S string, T string) string {
	for i := len(S) - 1; i >= 0; i-- {
		T = moveAllLetter(rune(S[i]), T)
	}
	return T
}

func moveAllLetter(char rune, word string) string {
	for i, letter := range word {
		if letter == char {
			word = string(letter) + word[:i] + word[i+1:]
		}
	}
	return word
}