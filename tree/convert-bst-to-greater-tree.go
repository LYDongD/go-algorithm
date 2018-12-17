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
func convertBST(root *TreeNode) *TreeNode {

	sum := 0
	convertBSTRecursively(root, &sum)
	return root
}

func convertBSTRecursively(root *TreeNode, sum *int) {

	if root == nil {
		return
	}

	convertBSTRecursively(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	convertBSTRecursively(root.Left, sum)

}

func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: -4}
	root.Left.Right = &TreeNode{Val: 1}
	convertBST(root)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val, root.Left.Left.Val, root.Left.Right.Val)
}
