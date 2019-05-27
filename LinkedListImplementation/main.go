package main

import (
	"fmt"
	list "./SingleLinkedList"
)

func main() {
	firstList := list.LinkedList{}
	firstList.Push(1)
	firstList.Push(3)
	firstList.Push(2)
	firstList.Push(4)
	firstList.Push(5)
	firstList.Push(8)
	firstList.Push(7)
	firstList.Push(6)
	
	head := firstList.Head
	for {
		if head == nil {
			break
		}
		fmt.Println(head.Val)
		head = head.Next
	}
	firstList.SortList()
	head = firstList.Head
	for {
		if head == nil {
			break
		}
		fmt.Println(head.Val)
		head = head.Next
	}
	// head = doubleLinkedList.Tail
	// for {
	// 	if head == nil {
	// 		break
	// 	}
	// 	fmt.Println(head.Val)
	// 	head = head.Prev
	// }

}