package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// getIntersectionNode 返回两个链表的相交节点，如果不相交则返回nil
func getIntersectionNode(headA *ListNode, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	// 指针pA和pB分别指向两个链表的头节点
	pA := headA
	pB := headB

	// 当两个指针不相等时，继续循环
	for pA != pB {
		// 如果pA到达链表末尾，将其重新指向headB
		if pA == nil {
			pA = headB
		} else {
			// 否则移动到下一个节点
			pA = pA.Next
		}

		// 如果pB到达链表末尾，将其重新指向headA
		if pB == nil {
			pB = headA
		} else {
			// 否则移动到下一个节点
			pB = pB.Next
		}
	}

	// 当两个指针相等时，返回它们指向的节点（即相交节点）
	return pA
}

func main() {
	// 创建两个链表，模拟相交的情况
	// A: a1 → a2
	//        ↘
	//        c1 → c2 → c3
	//        ↗
	// B: b1 → b2
	a1 := &ListNode{Val: 1}
	a2 := &ListNode{Val: 2}
	c1 := &ListNode{Val: 3}
	c2 := &ListNode{Val: 4}
	c3 := &ListNode{Val: 5}
	a1.Next = a2
	a2.Next = c1
	c1.Next = c2
	c2.Next = c3

	b1 := &ListNode{Val: 6}
	b2 := &ListNode{Val: 7}
	b1.Next = b2
	b2.Next = c1

	fun := getIntersectionNode(a1, b1)
	if fun != nil {
		fmt.Println(fun.Val)
	} else {
		fmt.Println("null.")
	}
}
