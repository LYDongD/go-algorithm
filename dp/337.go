package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	dic := make(map[*TreeNode]int)
	return robRecursively(root, dic)
}

func robRecursively(root *TreeNode, dic map[*TreeNode]int) int {
	if root == nil {
		return 0
	}

	if dic[root] > 0 {
		return dic[root]
	}

	leftSum, rightSum := 0, 0
	if root.Left != nil {
		leftSum += robRecursively(root.Left.Left, dic) + robRecursively(root.Left.Right, dic)
	}

	if root.Right != nil {
		rightSum += robRecursively(root.Right.Left, dic) + robRecursively(root.Right.Right, dic)
	}

	skipLevelSum := root.Val + leftSum + rightSum
	nextLevelSum := robRecursively(root.Left, dic) + robRecursively(root.Right, dic)
	result := max(skipLevelSum, nextLevelSum)
	dic[root] = result
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {

	root := &TreeNode{}
	root.Val = 3
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 1}
	fmt.Println(rob(root))
}
