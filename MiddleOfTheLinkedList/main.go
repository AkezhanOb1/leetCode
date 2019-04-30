package main

import (
	"fmt"
)

func main() {
	newNode2 := &listNode{
		Val: 4,
		Next: nil,
	}
	newNode1 := &listNode{
		Val: 3,
		Next: newNode2,
	}
	newNode := &listNode{
		Val: 2,
		Next: newNode1,
	}
	node := &listNode{
		Val: 1,
		Next: newNode,
	}

	fmt.Println(middleNode(node))

}

type listNode struct {
	     Val int
	     Next *listNode
}

func middleNode(head *listNode) *listNode {
	if head==nil||head.Next==nil{
        return head
    }
    slow:=head
    fast:=head
    for fast != nil && fast.Next!=nil{
        slow=slow.Next
        fast=fast.Next.Next
    }
    return slow
}