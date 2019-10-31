package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findContinuousNode(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return reuslt
	}

	findRecursively(root, []int{})
}

func findRecursively(node *TreeNode, current []int) []int {

	if node == nil {
		return current
	}
	
	if node.Left == nil && node.Right == nil {
		current = append(current, node.Val)
		return current
	}

	left := findRecursively(node.Left)
	right := findRecursively(node.Right)

	arr1, exist1 := parentChild(node, left, right)
	if exist1 {
		current = arr1
	}

	arr, exist2 := childParentChild((node, left, right) {
	if exist2 {
		current = arr
	}
	return current
}

func main() {
	fmt.Println("vim-go")
}
