package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	if root == nil {
		panic("tree is nil")
	}

	minVal := int(^uint(0) >> 1)
	inOrder(root, &minVal, nil)
	return minVal
}

func inOrder(node *TreeNode, minVal *int, prevNode *TreeNode) *TreeNode {
	if node == nil {
		return prevNode
	}

	prevNode = inOrder(node.Left, minVal, prevNode)

	if prevNode != nil {
		*minVal = min(*minVal, node.Val-prevNode.Val)
	}

	prevNode = node

	return inOrder(node.Right, minVal, prevNode)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	root := &TreeNode{34, nil, nil}
	root.Left = &TreeNode{27, nil, nil}
	root.Right = &TreeNode{50, nil, nil}
	root.Right.Left = &TreeNode{44, nil, nil}
	root.Right.Right = &TreeNode{58, nil, nil}
	fmt.Println(minDiffInBST(root))
}
