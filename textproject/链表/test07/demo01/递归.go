package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{}
	p := list
	n, r := 0, 0
	for l1 != nil || l2 != nil {
		n = l1.Val + l2.Val
		p.Val = n
		if n >= 10 {
			r = n / 10
			p.Val = r
		} else {

		}

		l1 = l1.Next
		l2 = l2.Next
		p = p.Next
	}
	return p.Next
}

func main() {
	headA := &ListNode{Val: 2}
	headA.Next = &ListNode{Val: 4}
	headA.Next.Next = &ListNode{Val: 3}

	headB := &ListNode{Val: 5}
	headB.Next = &ListNode{Val: 6}
	headB.Next.Next = &ListNode{Val: 4}
	list := addTwoNumbers(headA, headB)
	for list != nil {
		fmt.Print(list.Val, " ")
		list = list.Next
	}

}
