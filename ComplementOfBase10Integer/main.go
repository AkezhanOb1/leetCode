package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(bitwiseComplement(5))
}

func bitwiseComplement(N int) int {
	binaryN := strconv.FormatInt(int64(N), 2)
	reversedN := ""
	for _, bit := range binaryN {
		if bit == '0' {
			reversedN += "1"
		} else {
			reversedN += "0"
		}
	}
	res, _ := strconv.ParseInt(reversedN, 2, 64)
	return int(res)
}
