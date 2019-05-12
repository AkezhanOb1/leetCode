package list 

// DoubleLinkedList - represents a whole list at all, it has two filds
// Head is a pointer to the very begining of the list,
// Tail is a pointer to the very end of the list,
// if there is only one record in the list, head and tail 
// reference to the same node
type DoubleLinkedList struct {
	Head *Node
	Tail *Node
	Length int
}

// Node - represents single record in the Doubly Linked List
// it has two filds, actual value and the reference to the next node
// if there is no next node Next field will contain a nil value 
type Node struct {
	Prev *Node
	Next *Node 
	Val int
}


// Push - accepts an integer and puts it into the Doubly Linked List
// as the last item(tail)
func (l *DoubleLinkedList) Push(val int) {
	node := &Node{Val: val, Prev: l.Tail}
	if l.Head == nil {
		l.Head = node
	}else {
		l.Tail.Next = node
	}
	l.Tail = node
	l.Length++
}


// Pop - removes a last item from the Doubly Linked list
// and sets penultimate node as a tail
func (l *DoubleLinkedList) Pop() *Node {
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
	lastButOne := last.Prev
	lastButOne.Next = nil
	l.Tail = lastButOne
	l.Length--
	return last
}


// Shift - removes a node from the very begining 
// of the doubly linked list and sets head to the next element
// returns shifted node
func (l *DoubleLinkedList) Shift() *Node {
	head := l.Head
	if l.Length == 1 || l.Length == 0 {
		l.Head = nil
		l.Tail = nil
		l.Length = 0
		return head
	}
	l.Head = head.Next
	l.Head.Prev = nil
	head.Next = nil
	l.Length--
	return head
}


// Unshift - adds a node to the very begining 
// of the doubly linked list and sets head to the new element
// returns a new node
func (l *DoubleLinkedList) Unshift(val int) *Node {
	newHead := &Node {
		Val: val,
		Next: l.Head,
	}
	if l.Length == 0 {
		l.Head = newHead
		l.Tail = newHead
	}else {
		l.Head.Prev = newHead
		l.Head = newHead	
	}
	l.Length++
	return newHead
}

// Get - returns a node by its index
// if there is no node with that index 
// returns nil
func (l *DoubleLinkedList) Get(index int) *Node {
	if l.Length <= index || index < 0{
		return nil
	}
	if index < l.Length / 2 {
		node := l.Head
		for i := 0; i < index; i++ {
			node = node.Next
		}
		return node
	}

	node := l.Tail 
	index = l.Length - index
	for i := 0; i < index - 1; i++ {
		node = node.Prev
	}
	return node	
}


// Set - update the value of the node by its index
// and returns true if the node was updated
// otherwise returns false
func (l *DoubleLinkedList) Set(index int, val int) bool {
	node := l.Get(index)
	if node == nil {
		return false
	}
	node.Val = val
	return true
}



// Insert - inserts a new node into the doubly linked list
// by its position and returns true if it was inserted
// otherwise returns false
func (l *DoubleLinkedList) Insert(index int, val int) bool {
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
			Prev: node,
		}
		l.Length++
		node.Next = newNode
	}
	return true
}


// Remove - removes a node from the doubly linked list
// by its position and returns true if it was removed
// otherwise returns false
func (l *DoubleLinkedList) Remove(index int) bool {
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
		node.Next.Prev = node
		node.Next = nil
		node.Prev = nil
		l.Length--
	}
	return true
}