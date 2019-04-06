package main

import (
	"fmt"
)

func main() {
	fmt.Println(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	counter := 0
	if x < y {
		x, y = y, x
	}
	for ; x != 0; x, y = x>>1, y>>1 {
		if (x & 1) != (y & 1) {
			counter++
		}
	}
	return counter

}
