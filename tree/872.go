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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return false
	}

	seq1 := preOrder(root1, []int{})
	seq2 := preOrder(root2, []int{})
	return isSame(seq1, seq2)
}

func preOrder(root *TreeNode, seq []int) []int {
	if root.Left == nil && root.Right == nil {
		seq = append(seq, root.Val)
		return seq
	}

	if root.Left != nil {
		seq = preOrder(root.Left, seq)
	}

	if root.Right != nil {
		seq = preOrder(root.Right, seq)
	}
	return seq
}

func isSame(seq1 []int, seq2 []int) bool {
	if len(seq1) != len(seq2) {
		return false
	}

	for index, val := range seq1 {
		if val != seq2[index] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("")
}
