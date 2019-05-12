package main

import (
	"fmt"
	"strings"
	"sort"
)

func main(){
	fmt.Println(replaceWords([]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs"))
}

func replaceWords(dict []string, sentence string) string {
	sort.Strings(dict)
	splitedSentence := strings.Fields(sentence)
	for _, word := range dict {
		for i,successor := range splitedSentence {
			
			index := strings.HasPrefix(successor, word)
			if index == true {
				splitedSentence[i] = word
			}
		}
	}
	return strings.Join(splitedSentence, " ")
}

