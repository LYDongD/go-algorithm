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

    result,_,_ := isValid(root)
    return result
}

func isValid(node *TreeNode) (bool, int, int) {

    min, max := node.Val, node.Val
    if node.Left != nil {
	leftResult, leftMin, leftMax := isValid(node.Left)
	if !leftResult {
	    return false,0,0
	}

	if node.Val <= leftMax {
	    return false,0,0
	}

	min = leftMin
    }

    if node.Right != nil {
	rightResult, rightMin, rightMax := isValid(node.Right)
	if !rightResult {
	    return false,0,0
	}

	if node.Val >= rightMin {
	    return false,0,0
	}

	max = rightMax
    }
    fmt.Println(min, max)

    return true, min, max

}

func main() {
    root := &TreeNode{Val:1}
    root.Left = &TreeNode{Val:1}
    //root.Right = &TreeNode{Val:8}
    //root.Right.Left = &TreeNode{Val:7}
    //root.Right.Right = &TreeNode{Val:9}

    fmt.Println(isValidBST(root))
}
