package main

import "fmt"

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
func isSubtree(s *TreeNode, t *TreeNode) bool {

	if s == nil || t == nil {
		return false
	}

	return isSame(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(node1 *TreeNode, node2 *TreeNode) bool {

	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	if node1.Val == node2.Val {
		return isSame(node1.Left, node2.Left) && isSame(node1.Right, node2.Right)
	}

	return false
}

func main() {

	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	//	root.Left.Right.Left = &TreeNode{Val: 0}

	t := &TreeNode{Val: 4}
	t.Left = &TreeNode{Val: 1}
	t.Right = &TreeNode{Val: 2}

	fmt.Println(isSubtree(root, t))
}
