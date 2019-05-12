package  main

func main() {

}

// ListNode - represents singly linked list
type ListNode struct {
	Val int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
    res := []int{}
    headPointer := head 
    for headPointer != nil {
        head = headPointer.Next
        for head != nil {
            if head.Val > headPointer.Val {
                res = append(res, head.Val)
                break
            }
            head = head.Next
        }
        if head == nil {
            res = append(res, 0)
        }
        headPointer = headPointer.Next
    }
    return res
}