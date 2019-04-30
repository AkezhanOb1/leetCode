package main

import (
	"fmt"
)

func main(){
	fmt.Println(findDuplicates([]int{4,3,2,7,8,2,3,1}))
}

func findDuplicates(nums []int) []int {
	counter := make([]int, len(nums) + 1)
    res := []int{}
    for _, val := range nums {
        counter[val]++
    }
    for i, v := range counter {
        if v > 1 {
            res = append(res, i)
        }
    }
    return res
}