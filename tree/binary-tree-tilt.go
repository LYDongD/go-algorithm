package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	_, tilt := findSumAndTilt(root)
	return tilt
}

func findSumAndTilt(root *TreeNode) (sum, tilt int) {
	if root == nil {
		return 0, 0
	}

	ls, lt := findSumAndTilt(root.Left)
	rs, rt := findSumAndTilt(root.Right)
	return ls + rs + root.Val, lt + rt + abs(ls, rs)
}

func abs(a, b int) int {
	result := a - b
	if result < 0 {
		result = -result
	}

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	root.Left = left
	root.Right = right

	fmt.Println(findTilt(root))
}
