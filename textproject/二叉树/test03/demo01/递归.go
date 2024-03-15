package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	fmt.Println(invertTree(root))
}
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	Left := invertTree(root.Left)
	Right := invertTree(root.Right)
	root.Left = Right
	root.Right = Left
	return root
}
