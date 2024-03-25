package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
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
	for list != nil {
		fmt.Print(list.Val, " ")
		list = list.Next
	}

}
