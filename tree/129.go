package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var sum int

func sumNumbers(root *TreeNode) int {
    if root == nil {
	return 0
    }

    sum = 0
    preorderCount(root, 0)
    return sum
}

func preorderCount(node *TreeNode, rootToNodeSum int) {
    if node.Left == nil && node.Right == nil  {
	rootToNodeSum = rootToNodeSum * 10 + node.Val
	sum += rootToNodeSum
	return
    }

    rootToNodeSum = rootToNodeSum * 10 + node.Val
    if node.Left != nil {
	preorderCount(node.Left, rootToNodeSum)
    }

    if node.Right != nil {
	preorderCount(node.Right, rootToNodeSum)
    }
}

func main() {
    root := &TreeNode{Val:1}
    root.Left = &TreeNode{Val:2}
    root.Right = &TreeNode{Val:3}

    fmt.Println(sumNumbers(root))
}

