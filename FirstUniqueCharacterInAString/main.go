package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("leetcodeleetcode"))
}

func firstUniqChar(s string) int {
	dic := counter(s)
	for i, letter := range s {
		if val, ok := dic[letter]; ok {
			if val == 1 {
				return i
			}
		}
	}
	return -1
}

func counter(s string) (counter map[rune]int) {
	counter = make(map[rune]int)
	for _, letter := range s {
		if _, ok := counter[letter]; ok {
			counter[letter]++
		} else {
			counter[letter] = 1
		}
	}

	return
}
