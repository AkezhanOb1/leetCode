package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numJewelsInStones("z", "ZZ"))
}

func numJewelsInStones(J string, S string) int {
	var counter int
	for _, i := range S {
		counter += strings.Count(J, string(i))
	}
	return counter
}
