package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func main() {
	head := &ListNode{Val: 3}
	node02 := &ListNode{Val: 2}
	node03 := &ListNode{Val: 0}
	node04 := &ListNode{Val: -4}
	head.Next = node02
	node02.Next = node03
	node03.Next = node04

	node04.Next = node02
	h := hasCycle(head)
	fmt.Println(h)
	//取消环形链表
	node04.Next = nil

	h = hasCycle(head)
	fmt.Println(h)
}
