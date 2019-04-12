package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(checkAnagram("aqua", "aqwe"))
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func checkAnagram(firstString string, secondString string) bool {
	if firstString == "" && secondString == "" {
		return true
	}

	if len(firstString) != len(secondString) {
		return false
	}

	counter := make(map[rune]int)
	for _, v := range firstString {
		if _, ok := counter[v]; ok {
			counter[v]++
			continue
		}
		counter[v] = 1
	}
	for _, v := range secondString {
		if _, ok := counter[v]; ok {
			counter[v]--
			continue
		}
		return false
	}

	return true
}
