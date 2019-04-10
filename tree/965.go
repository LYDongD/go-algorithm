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
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
    if root == nil {
	return true
    }

    return preOrder(root, root.Val)
}

func preOrder(node *TreeNode, val int) bool {
    if node == nil {
	return true
    }

    if node.Val != val {
	return false
    }

    result := preOrder(node.Left, val)
    if result {
	result = preOrder(node.Right, val)
    }

    return result
}



