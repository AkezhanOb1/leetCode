package main

import "fmt"

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20}))
}

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0

	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five++
			continue
		}
		if bills[i] == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
			continue
		}
		if bills[i] == 20 {
			if ten == 0 {
				if five < 3 {
					return false
				}
				five -= 3
				continue
			}
			if five == 0 {
				return false
			}
			ten--
			five--
		}
	}
	return true
}
