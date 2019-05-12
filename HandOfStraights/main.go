package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isNStraightHand([]int{2,3,4,5,6,7,8 }, 3))
}

func isNStraightHand(hand []int, W int) bool {
	sort.Ints(hand)
	for i := 0; i < W; i++ {
		arr, sortArr := fillArr(hand, W)
		if len(arr) != W {
			return false
		}
		hand = sortArr
	}
	return true
}

func fillArr(arr []int, size int) ([]int, []int){
	fmt.Println("START", arr)
	res := []int{}
	slicedArr := arr
	lasEl := -999
	i := 0
	for len(res) == size {
		if len(res) == size {
			return res, slicedArr
		}
		if lasEl != slicedArr[i] {
			lasEl = slicedArr[i]
			res = append(res, slicedArr[i])
			fmt.Println("lasEl",lasEl)
			fmt.Println("sliceArr", slicedArr[:i], slicedArr[i+1:])
			fmt.Println("RES", res)
			slicedArr = append(slicedArr[:i], slicedArr[i+1:]...)
			fmt.Println("SLICED", slicedArr)
		}
	}
	return res, slicedArr
} 