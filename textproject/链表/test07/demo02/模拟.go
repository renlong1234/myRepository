package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	tail := &ListNode{}
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
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
