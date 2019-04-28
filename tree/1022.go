package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	preOrder(root, 0, &sum)
	return sum
}

func preOrder(node *TreeNode, num int, sum *int) {
	if node == nil {
		return
	}

	num = num*2 + node.Val
	if node.Left == nil && node.Right == nil {
		*sum = *sum + num
	} else {
		preOrder(node.Left, num, sum)
		preOrder(node.Right, num, sum)
	}
}

func main() {
	fmt.Println("vim-go")
}
