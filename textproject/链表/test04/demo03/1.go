package main

import (
	"fmt"
)

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle 检测链表是否有环
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
	//// 创建链表
	//head := &ListNode{Val: 1}
	//node2 := &ListNode{Val: 2}
	//node3 := &ListNode{Val: 3}
	//head.Next = node2
	//node2.Next = node3
	//
	//// 创建一个环形链表
	//node3.Next = node2
	//
	//// 检测环形链表
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next.Next = head.Next
	result := hasCycle(head)

	fmt.Println("是否存在环形链表:", result) // 输出: 是否存在环形链表: true

	// 打破环
	head.Next.Next.Next.Next = nil

	// 再次检测环形链表
	result = hasCycle(head)
	fmt.Println("是否存在环形链表:", result) // 输出: 是否存在环形链表: false
}
