package main

import (
	"fmt"
)

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}

func findWords(words []string) []string {
	firstRow := map[rune]bool{
		'q': true,
		'w': true,
		'e': true,
		'r': true,
		't': true,
		'y': true,
		'u': true,
		'i': true,
		'o': true,
		'p': true,
		'Q': true,
		'W': true,
		'E': true,
		'R': true,
		'T': true,
		'Y': true,
		'U': true,
		'I': true,
		'O': true,
		'P': true,
	}
	secondRow := map[rune]bool{
		'a': true,
		's': true,
		'd': true,
		'f': true,
		'g': true,
		'h': true,
		'j': true,
		'k': true,
		'l': true,
		'A': true,
		'S': true,
		'D': true,
		'F': true,
		'G': true,
		'H': true,
		'J': true,
		'K': true,
		'L': true,
	}
	thirdRow := map[rune]bool{
		'z': true,
		'x': true,
		'c': true,
		'v': true,
		'b': true,
		'n': true,
		'm': true,
		'Z': true,
		'X': true,
		'C': true,
		'V': true,
		'B': true,
		'N': true,
		'M': true,
	}
	res := []string{}
	for _, word := range words {
		if _, ok := firstRow[rune(word[0])]; ok {
			checked := checker(word[1:], firstRow)
			if checked == true {
				res = append(res, word)
			}
		} else if _, ok := secondRow[rune(word[0])]; ok {
			checked := checker(word[1:], secondRow)
			if checked == true {
				res = append(res, word)
			}
		} else {
			checked := checker(word[1:], thirdRow)
			if checked == true {
				res = append(res, word)
			}
		}
	}
	return res
}

func checker(word string, maper map[rune]bool) bool {
	for _, val := range word {
		if _, ok := maper[val]; !ok {
			return false
		}
	}
	return true
}
