package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(queryString("0110", 4))
}


func queryString(S string, N int) bool {

	for i := 1; i <= N; i++ {
		if strings.Index(S, strconv.FormatInt(int64(i), 2)) == -1 {
			return false
		}
	}
	
	return true
}