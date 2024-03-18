package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var p1 *ListNode
	p2 := head
	for p2 != nil {
		nextTemp := p2.Next
		p1, p2.Next = p2, p1
		p2 = nextTemp
	}
	return p1
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " 	")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// 创建链表：1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Println("原始链表：")
	printList(head)

	newHead := reverseList(head)
	fmt.Println("反转后的链表：")
	printList(newHead)
}
