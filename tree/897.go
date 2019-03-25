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

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	dummy := &TreeNode{0, nil, nil}
	cursor := dummy
	inOrder(root, cursor)
	return dummy.Right
}

func inOrder(node *TreeNode, cursor *TreeNode) *TreeNode {
	if node == nil {
		return cursor
	}

	cursor = inOrder(node.Left, cursor)
	cursor.Left = nil
	cursor.Right = node
	cursor = node
	cursor = inOrder(node.Right, cursor)
	return cursor
}

func main() {
	root := &TreeNode{379, &TreeNode{826, nil, nil}, nil}
	node := increasingBST(root)
	cursor := node
	for cursor != nil {
		fmt.Println(cursor.Val)
		cursor = cursor.Right
	}
}
