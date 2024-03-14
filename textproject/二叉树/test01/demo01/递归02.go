package main

import (
	"fmt"
)

// TreeNode represents a node in a binary tree
type TreeNode02 struct {
	Val   int
	Left  *TreeNode02
	Right *TreeNode02
}

// inorderTraversalRecursive performs an in-order traversal recursively
func inorderTraversalRecursive(root *TreeNode02) {
	if root == nil {
		return
	}
	inorderTraversalRecursive(root.Left)
	fmt.Println(root.Val)
	inorderTraversalRecursive(root.Right)
}

func main() {
	// Create a binary tree
	//      1
	//     / \
	//    2   3
	//   / \
	//  4   5
	root := &TreeNode02{
		Val: 1,
		Left: &TreeNode02{
			Val: 2,
			Left: &TreeNode02{
				Val: 4,
			},
			Right: &TreeNode02{
				Val: 5,
			},
		},
		Right: &TreeNode02{
			Val: 3,
		},
	}
	// Perform in-order traversal
	fmt.Println("该二叉树的中序遍历为：")
	inorderTraversalRecursive(root)
}
