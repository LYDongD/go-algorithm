package array

import "fmt"
import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	currentVal := strconv.Itoa(t.Val)

	if t.Left == nil && t.Right == nil {
		return currentVal
	}

	if t.Left == nil {
		return currentVal + "()" + "(" + tree2str(t.Right) + ")"
	}

	if t.Right == nil {
		return currentVal + "(" + tree2str(t.Left) + ")"
	}

	return currentVal + "(" + tree2str(t.Left) + ")" + "(" + tree2str(t.Right) + ")"
}

func array() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	fmt.Println(tree2str(root))
}
