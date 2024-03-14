package main

import (
	"fmt"
)

// TreeNode02 represents a node in a binary tree
type TreeNode02 struct {
	Val   int
	Left  *TreeNode02
	Right *TreeNode02
}

// maxDepth02 recursively calculates the maximum depth of a binary tree
func maxDepth02(root *TreeNode02) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth02(root.Left)
	rightDepth := maxDepth02(root.Right)

	// The depth of the tree is the maximum depth of its left and right subtrees, plus one for the root node
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func main() {
	// Create a binary tree
	//      1
	//     / \
	//    2   3
	//   / \   \
	//  4   5   6
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
			Right: &TreeNode02{
				Val: 6,
			},
		},
	}

	// Calculate the depth of the binary tree
	depth := maxDepth02(root)
	fmt.Printf("Depth of the binary tree is: %d\n", depth)
}
