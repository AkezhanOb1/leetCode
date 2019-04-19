package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
}

func toGoatLatin(S string) string {
	resString := ""
	aAdder := "a"
	splittedArr := strings.Fields(S)
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	for i, val := range splittedArr {
		if _, ok := vowels[rune(val[0])]; ok {
			resString = resString + splittedArr[i] + "ma" + aAdder
		} else {
			reOrdered := val[1:] + string(val[0])
			resString = resString + reOrdered + "ma" + aAdder
		}
		if i != len(splittedArr)-1 {
			resString += " "
		}

		aAdder += "a"
	}

	return resString
}
