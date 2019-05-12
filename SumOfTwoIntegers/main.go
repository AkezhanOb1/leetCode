
package main

import "fmt"

func main() {
	fmt.Println(getSum(-20, 3))
}


func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	return getSum(2*(a&b), a^b)
}
