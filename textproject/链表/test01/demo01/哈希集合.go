package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a1 := &ListNode{Val: 4}
	a2 := &ListNode{Val: 1}
	c1 := &ListNode{Val: 8}
	c2 := &ListNode{Val: 4}
	c3 := &ListNode{Val: 5}
	a1.Next = a2
	a2.Next = c1
	c1.Next = c2
	c2.Next = c3
	b1 := &ListNode{Val: 5}
	b2 := &ListNode{Val: 6}
	b3 := &ListNode{Val: 1}
	b1.Next = b2
	b2.Next = b3
	b3.Next = c1
	Count := getIntersectionNode(a1, b1)
	if Count != nil {
		fmt.Println(Count.Val)
	} else {
		fmt.Println("null")
	}
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}
