package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
    if root == nil {
	return true
    }

    MIN := ^(int(^uint(0) >> 1))
    return inOrder(root, &MIN)
}

func inOrder(node *TreeNode, lastVal *int) bool {

    if node == nil {
	return true
    }

    result := inOrder(node.Left, lastVal)
    if result == false {
	return false
    }

    if node.Val <= *lastVal {
	return false
    }

    *lastVal = node.Val

    return inOrder(node.Right, lastVal)
}

func main() {
    root := &TreeNode{Val:1}
    root.Left = &TreeNode{Val:1}
    //root.Right = &TreeNode{Val:8}
    //root.Right.Left = &TreeNode{Val:7}
    //root.Right.Right = &TreeNode{Val:9}

    fmt.Println(isValidBST(root))
}

