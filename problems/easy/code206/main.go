package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := new(ListNode)
	reverse(head, head.Next, newHead)
	head.Next = nil
	return newHead.Next
}

func reverse(parent, child, newHead *ListNode) {
	if child == nil {
		newHead.Next = parent
		return
	}
	reverse(child, child.Next, newHead)
	child.Next = parent
}

func main() {
	one := &ListNode{
		Val: 1,
	}
	two := &ListNode{
		Val: 2,
	}
	//three := &ListNode{
	//	Val: 3,
	//}
	//four := &ListNode{
	//	Val: 4,
	//}
	one.Next = two
	//two.Next = three
	//three.Next = four

	printListNode(reverseList(one))

}

func printListNode(head *ListNode) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
}
