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
	return findTiltRecursively(root, 0)
}

func findTiltRecursively(root *TreeNode, tilt int) int {
	if root == nil {
		return tilt
	}

	tilt += calculateTilt(root)
	tilt = findTiltRecursively(root.Left, tilt)
	tilt = findTiltRecursively(root.Right, tilt)

	return tilt
}

func calculateTilt(node *TreeNode) int {
	diff := calSum(node.Left, 0) - calSum(node.Right, 0)
	if diff > 0 {
		return diff
	} else {
		return -diff
	}
}

func calSum(node *TreeNode, sum int) int {
	if node == nil {
		return sum
	}

	sum += node.Val
	sum = calSum(node.Left, sum)
	sum = calSum(node.Right, sum)

	return sum
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
