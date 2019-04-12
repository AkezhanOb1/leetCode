package main

import (
	"fmt"
)

func main() {
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
}

func commonChars(A []string) []string {
	counter := make([][26]int, len(A))
	min := 0
	result := make([]string, 0)
	for i, v := range A {
		for _, c := range v {
			counter[i][c-'a']++
		}
	}
	for i := 0; i < 26; i++ {
		min = counter[0][i]
		for j := 1; j < len(counter); j++ {
			if counter[j][i] < min {
				min = counter[j][i]
			}
		}
		for j := 0; j < min; j++ {
			result = append(result, string('a'+i))
		}
	}
	return result
}
