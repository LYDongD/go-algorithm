package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	path := []int{}
	return findPath(root, 0, sum, path, result)
}

func findPath(root *TreeNode, currentSum, sum int, path []int, result [][]int) [][]int {
	if root == nil && currentSum == sum {
		result = append(result, path)
		return result
	}

	currentSum += root.Val
	if currentSum > sum {
		return result
	}

	if path == nil {
		path = []int{}
	}
	path = append(path, root.Val)
	result = findPath(root.Left, currentSum, sum, path, result)
	result = findPath(root.Right, currentSum, sum, path, result)
	return result
}

func main() {
	fmt.Println("vim-go")
}
