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
func diameterOfBinaryTree(root *TreeNode) int {

	count := 1
	nodeCount(root, &count)
	return count - 1
}

func nodeCount(node *TreeNode, count *int) int {
	if node == nil {
		return 0
	}

	leftCount := nodeCount(node.Left, count)
	rightCount := nodeCount(node.Right, count)

	if *count < leftCount+rightCount+1 {
		*count = leftCount + rightCount + 1
	}

	if leftCount > rightCount {
		return leftCount + 1
	} else {
		return rightCount + 1
	}
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Println(diameterOfBinaryTree(root))
}
