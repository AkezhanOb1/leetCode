package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(isAnagram("anagram", "nagaraqm"))
}

func isAnagram(s string, t string) bool {
	var firstWordCounter map[rune]int
	var secondWordCounter map[rune]int
	firstWordCounter = letterCounter(s)
	secondWordCounter = letterCounter(t)
	return reflect.DeepEqual(firstWordCounter, secondWordCounter)
}

func letterCounter(s string) map[rune]int {
	counter := make(map[rune]int)
	for _, letter := range s {
		if _, ok := counter[letter]; ok {
			counter[letter]++
		} else {
			counter[letter] = 1
		}
	}
	return counter
}
