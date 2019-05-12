package main

import (
	"fmt"
	list "./SingleLinkedList"
)

func main() {
	firstList := list.LinkedList{}
	firstList.Push(1)
	firstList.Push(2)
	firstList.Push(4)
	
	secondList := list.LinkedList{}
	secondList.Push(1)
	secondList.Push(3)
	secondList.Push(4)

	newList := list.MergeTwoLists(firstList, secondList)
	head := newList.Head
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