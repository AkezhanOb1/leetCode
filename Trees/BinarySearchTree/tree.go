package tree

// Node - represents a single node in the binary search tree
// it has 3 properties: Val - actual value that is stores,
// link to the left node which should be less than current node,
// link to the right node which should be more than current node.
type Node struct {
	Val int
	Left *Node
	Right *Node
}

// BinarySearchTree - represents a BTS and contains only one properties
// link to the root(very first) node
type BinarySearchTree struct {
	Root *Node
}


// Constructor - defines a root node int the BTS
// accepts an integer to be stored as value of the root node
func (b *BinarySearchTree) Constructor(val int) {
	if ok := b.CheckRoot(); !ok {
		b.Root = &Node {
			Val: val,
		}
	}
}

// CheckRoot - checks if BTS has a root node
// if there is no root returns false, otherwise returns true
func (b *BinarySearchTree) CheckRoot() bool{
	if b.Root != nil {
		return true
	}
	return false
}

// Insert - inserts a new node into the BTS, if the val is less than
// root's val, the it will be inserted on the left half, otherwise on the right 
func (b *BinarySearchTree) Insert(val int) {
	newNode := &Node {
		Val: val,
	}
	if ok := b.CheckRoot(); ok {
		node := b.Root
		for node != nil {
			if node.Val > val {
				if node.Left == nil {
					node.Left = newNode
					break
				}
				node = node.Left
			}else if node.Val < val {
				if node.Right == nil {
					node.Right = newNode
					break
				}
				node = node.Right
			}else {
				break
			}
		}
	}else {
		b.Root = newNode
	}
}

// Find - searches for a node with a specific value, the value is sent 
// as an argument of the function, if BTS has a node with sush value 
// it returns it, if there is no such Node the function returns nil
func (b *BinarySearchTree) Find(val int) *Node {
	node := b.Root
	for node != nil {
		if node.Val < val {
			node = node.Right
		}else if node.Val > val {
			node = node.Left
		}else {
			return node
		}
	}
	return nil
}

// BFS - is a one of the many tree traversal methods 
// BFS - stands for Breadth First Search
// traversing from the very begining to the very end 
// from left to right  and returns 
// every value in the tree
func (b *BinarySearchTree) BFS() []int {
	var queue []*Node
	var data  []int
	queue = append(queue, b.Root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		data = append(data, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return data
}

// DFS - is a another one way of traversing the tree
// DFS - stands for Depth First Search
// the concept idea is to travers from the very begining
// to the very end and first left halve, after that
// right halve traversal, returns every value in the tree
func (b *BinarySearchTree) DFS() []int {
	v := val{
		data: []int{},
	}
	node := b.Root
	dfsHelper(node, &v)
	return v.data
}

type val struct {
	data []int
}

func dfsHelper(node *Node, v *val) {
	v.data = append(v.data, node.Val)
	if node.Left != nil {
		dfsHelper(node.Left, v)
	} 
	if node.Right != nil {
		dfsHelper(node.Right, v)
	}
}

// FindDepth - returns depth of the entire tree
func (b *BinarySearchTree) FindDepth() int {
	return findDepthHelper(b.Root)
}

func findDepthHelper(node *Node) int{
	if node == nil {
		return 0
	}
	maxRight := findDepthHelper(node.Right)
	maxLeft := findDepthHelper(node.Left)
	if maxLeft > maxRight {
		return maxLeft + 1
	}
	return maxRight + 1
}