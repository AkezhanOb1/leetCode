package list

import (
	"sort"
)

// Node - represents single record in the Linked List
// it has two filds, actual value and the reference to the next node
// if there is no next node Next field will contain a nil value 
type Node struct {
	Val int
	Next *Node
}

// LinkedList - represents a whole list at all, it has two filds
// Head is a pointer to the very begining of the list,
// Tail is a pointer to the very end of the list,
// if there is only one record in the list, head and tail 
// reference to the same node
type LinkedList struct {
	Head *Node
	Tail *Node
	Length int
}

// Push - accepts an integer and puts it into the Linked List
// as the last item(tail)
func (l *LinkedList) Push(val int) {
	node := &Node{Val: val}
	if l.Head == nil {
		l.Head = node
	}else {
		l.Tail.Next = node
	}
	l.Tail = node
	l.Length++
}

// Pop - removes a last item from the linked list
// and sets penultimate node as a tail
func (l *LinkedList) Pop() *Node {
	last := l.Tail
	len := l.Length
	if len == 0 {
		return last
	}else if len == 1 {
		l.Length = 0
		l.Head = nil
		l.Tail = nil
		return last
	}
	lastButOne := l.Head
	for i := 1; i < len - 1; i++ {
		lastButOne = lastButOne.Next
	}
	lastButOne.Next = nil
	l.Tail = lastButOne

	l.Length--
	return last
}

// Shift - removes a node from the very begining 
// of the linked list and sets head to the next element
// returns shifted node
func (l *LinkedList) Shift() *Node {
	head := l.Head
	if l.Length == 1 || l.Length == 0 {
		l.Head = nil
		l.Tail = nil
		l.Length = 0
		return head
	}
	l.Head = l.Head.Next
	l.Length--
	return head
}

// Unshift - adds a node to the very begining 
// of the linked list and sets head to the new element
// returns a new node
func (l *LinkedList) Unshift(val int) *Node {
	lastHead := l.Head
	newHead := &Node {
		Val: val,
		Next: lastHead,
	}
	l.Head = newHead
	if l.Length == 0 {
		l.Tail = newHead
	}
	l.Length++
	return newHead
}

// Get - returns a node by its index
// if there is no node with that index 
// returns nil
func (l *LinkedList) Get(index int) *Node {
	if l.Length <= index || index < 0{
		return nil
	}
	node := l.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node
}

// Set - update the value of the node by its index
// and returns true if the node was updated
// otherwise returns false
func (l *LinkedList) Set(index int, val int) bool {
	node := l.Get(index)
	if node == nil {
		return false
	}
	node.Val = val
	return true
}

// Insert - inserts a new node into the linked list
// by its position and returns true if it was inserted
// otherwise returns false
func (l *LinkedList) Insert(index int, val int) bool {
	if index < 0  || index > l.Length {
		return false
	}
	switch index {
	case 0: 
		l.Unshift(val)
	case l.Length: 
		l.Push(val)
	default: 
		node := l.Get(index - 1)
		newNode := &Node{
			Val: val,
			Next: node.Next,
		}
		l.Length++
		node.Next = newNode
	}
	return true
}


// Remove - removes a node from the linked list
// by its position and returns true if it was removed
// otherwise returns false
func (l *LinkedList) Remove(index int) bool {
	if index < 0  || index >= l.Length {
		return false
	}
	switch index {
	case 0: 
		l.Shift()
	case l.Length - 1: 
		l.Pop()
	default: 
		node := l.Get(index - 1)
		node.Next = node.Next.Next
		l.Length--
	}
	return true
}

// Reverse - changes the entire linked list
// in reverse order, head becomes a tail,
// tail becomes a head, so it looks like this
// 1->2->3 becomes 3->2->1
// 1<-2 3
func (l *LinkedList) Reverse() *LinkedList{
	node := l.Head
	l.Head = l.Tail
	l.Tail = node
	var next *Node
	var prev *Node
	for i := 0; i < l.Length; i++ {
		next = node.Next
		node.Next = prev
		prev = node
		node = next
	}
	l.Pop()
	return l
}

//SortList - sorts linked integer linked lists with time complexity n*Log(n)
func (l *LinkedList) SortList() {
    res := []int{}
    head := l.Head
    for head != nil {
        res = append(res, head.Val)
        head = head.Next
    }
    head = l.Head
    counter := 0
    sort.Ints(res)
    for head != nil {
        head.Val = res[counter]
        counter++
        head = head.Next
    }
}




// First - returns a head of the linkedList 
func (l *LinkedList) First() *Node {
	return l.Head
}

// Len - returns a lenth of the linkedList 
func (l *LinkedList) Len() int {
	return l.Length
}

// NextNode - returns a next node in the linked List
func (n *Node) NextNode() *Node{
	return n.Next
}

