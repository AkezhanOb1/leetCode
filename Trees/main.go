package main

import (
	"fmt"
	"./BinarySearchTree"
)

func main() {
	tr := tree.BinarySearchTree{}
	tr.Insert(10)
	tr.Insert(11)
	tr.Insert(12)
	tr.Insert(13)
	tr.Insert(14)
	tr.Insert(15)


	
	
	fmt.Println(tr.FindDepth())
}
