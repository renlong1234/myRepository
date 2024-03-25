package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	flag := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			flag = true
			break
		}
	}

	if !flag {
		return nil
	}

	p := head
	for p != slow {
		p = p.Next
		slow = slow.Next
	}

	return p
}

func main() {
	// 0 -> 1 -> 2 -> 3
	//                -> 2
	head := &ListNode{Val: 0}
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3

	node3.Next = node2
	entryNode := detectCycle(head)
	if entryNode != nil {
		fmt.Println("value=", detectCycle(head).Val)
	} else {
		fmt.Println("NULL")
	}
}
