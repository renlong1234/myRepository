package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list := &ListNode{}
	p := list
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	} else {
		p.Next = list2
	}
	return list.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}
func main() {
	headA := &ListNode{Val: 1}
	nodeA2 := &ListNode{Val: 2}
	nodeA3 := &ListNode{Val: 4}
	nodeA4 := &ListNode{Val: 5}
	nodeA5 := &ListNode{Val: 6}
	headA.Next = nodeA2
	nodeA2.Next = nodeA3
	nodeA3.Next = nodeA4
	nodeA4.Next = nodeA5

	headB := &ListNode{Val: 1}
	nodeB2 := &ListNode{Val: 3}
	nodeB3 := &ListNode{Val: 4}
	headB.Next = nodeB2
	nodeB2.Next = nodeB3
	list := mergeTwoLists(headA, headB)

	printList(list)

}
