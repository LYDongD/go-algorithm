package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}

	dic := make(map[int]int)
	for i, val := range inorder {
		dic[val] = i
	}

	return buildRecursively(dic, preorder, 0, 0, len(inorder)-1)
}

func buildRecursively(dic map[int]int, preorder []int, rootIndex, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	rootInOrderIndex := dic[preorder[rootIndex]]
	leftNodeIndex := rootIndex + 1
	rightNodeIndex := rootIndex + rootInOrderIndex - left + 1
	root := &TreeNode{Val: preorder[rootIndex]}
	root.Left = buildRecursively(dic, preorder, leftNodeIndex, left, rootInOrderIndex-1)
	root.Right = buildRecursively(dic, preorder, rightNodeIndex, rootInOrderIndex+1, right)
	return root
}

func main() {
	fmt.Println("vim-go")
}
