package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
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

	fmt.Println(detectCycle(head).Val)
}
