package main

import (
	"fmt"
)

func main() {
	a := []byte{'H','a','n','n','a','h'}
	reverseString(a)
}

func reverseString(s []byte)  {
	if len(s) == 0 {
		return 
	}
	reverseString(s[1:])
	fmt.Print(string(s[0]))
	
}