package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	height := countHeight(root)
	rightSubTreeHeight := countHeight(root.Right)
	if height == rightSubTreeHeight+1 {
		return 1<<uint(height) + countNodes(root.Right)
	} else {
		return 1<<uint((height-1)) + countNodes(root.Left)
	}
}

func countHeight(root *TreeNode) int {
	node := root
	height := -1
	for node != nil {
		height++
		node = node.Left
	}

	return height
}

func main() {

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	fmt.Println(countHeight(root))
	fmt.Println(countNodes(root))
}
