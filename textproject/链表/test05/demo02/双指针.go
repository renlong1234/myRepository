package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//	func detectCycle(head *ListNode) *ListNode {
//		low := head
//		fast := head.Next
//		for low < fast {
//			if head == nil || head.Next == nil {
//				return nil
//			}
//			if fast == nil || fast.Next == nil {
//				return nil
//			}
//			fast = fast.Next
//			low = low.Next
//		}
//		return n
//	}
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
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
