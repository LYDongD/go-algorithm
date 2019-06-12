package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func flatten(root *TreeNode)  {
    if root == nil {
	return
    }

    flatternRecursively(root)
}

func flatternRecursively(node *TreeNode) *TreeNode {
    if node.Left == nil && node.Right == nil {
	return node
    }

    var tail *TreeNode
    if node.Left != nil {
	tail = flatternRecursively(node.Left)
    }

    if tail != nil {
	tail.Right = node.Right
	tail.Left = nil
    }

    if node.Right != nil {
	tail = flatternRecursively(node.Right)
    }

    node.Right = node.Left
    node.Left = nil

    return tail
}

func main() {
    fmt.Println("haha")
}
