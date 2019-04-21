package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(addDigits(3812312311123123143))
	end := time.Since(start)
	fmt.Println(end)
}

func addDigits(num int) int {
	strNum := strconv.Itoa(num)
	splNum := strings.Split(strNum, "")
	res := num
	for len(splNum) > 1 {
		counter := 0
		for _, val := range splNum {
			i, _ := strconv.Atoi(val)
			counter += i
		}
		res = counter
		strNum = strconv.Itoa(res)
		splNum = strings.Split(strNum, "")
	}
	return res
}
