package main

import "fmt"

var ancestor *TreeNode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
	    return root
	}

	findPath(root, p, q)
	return ancestor
}

func findPath(root, p, q *TreeNode) bool {

    if root == nil {
	return false
    }

    mid := 0
    if root == p || root == q{
	mid = 1
    }

    left := 0
    if findPath(root.Left, p, q) {
	left = 1
    }

    right := 0
    if findPath(root.Right, p, q) {
	right = 1
    }

    if left + mid + right >= 2 {
	ancestor = root
    }

    return right + left + mid > 0
}

func main() {
    fmt.Println("haha")
}
