package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	var n int
	fmt.Print("请输入：n=")
	fmt.Scanln(&n)
	List := removeNthFromEnd(head, n)

	for List != nil {
		fmt.Print(List.Val, " ")
		List = List.Next
	}
}
