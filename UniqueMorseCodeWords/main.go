package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"rwjje", "aittjje", "auyyn", "lqtktn", "lmjwn"}))
}

func uniqueMorseRepresentations(words []string) int {
	var morseArr = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	var uniq = map[string]bool{}
	for _, v := range words {
		a := ""
		for _, code := range v {
			a += morseArr[code-97]
		}
		uniq[a] = true
	}
	return len(uniq)
}
