package main 

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println(countBits(5))
}

func countBits(num int) []int {
	res := []int{}
	for i := 0; i <= num; i++ {
		binar := strconv.FormatInt(int64(i), 2)
		binar = strings.Replace(binar, "0", "", -1)
		res = append(res, len(binar))
	}
	return res	
}