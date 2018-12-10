package main

import (
	"fmt"
)

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
func findMode(root *TreeNode) []int {

	countMap := make(map[int]int)

	countModeInPreOrder(root, countMap)

	fmt.Println("countMap: ", countMap)
	var result []int
	maxCount := -1
	for val, count := range countMap {
		if count > maxCount {
			maxCount = count
			result = nil
			result = append(result, val)
		} else if count == maxCount {
			result = append(result, val)
		}
	}

	return result
}

func countModeInPreOrder(root *TreeNode, countMap map[int]int) {

	if root == nil {
		return
	}

	countMap[root.Val]++

	countModeInPreOrder(root.Left, countMap)
	countModeInPreOrder(root.Right, countMap)

}

func main() {

	root := &TreeNode{Val: 1}
	right := &TreeNode{Val: 2}
	left := &TreeNode{Val: 2}

	right.Left = left
	root.Right = right

	fmt.Println(findMode(root))

}
