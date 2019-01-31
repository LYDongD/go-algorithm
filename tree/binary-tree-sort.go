package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sort(nums []int) []int {

	//构建二叉搜索树
	root := buidBinaryTree(nums)
	var result []int
	//中序遍历
	return inOrder(root, result)
}

func buidBinaryTree(nums []int) *TreeNode {

	var root *TreeNode
	for _, num := range nums {
		root = insertTree(root, num)
	}

	return root
}

func insertTree(root *TreeNode, num int) *TreeNode {

	if root == nil {
		root = &TreeNode{num, nil, nil}
		return root
	}

	if root.Val > num {
		root.Left = insertTree(root.Left, num)
	} else if root.Val < num {
		root.Right = insertTree(root.Right, num)
	}

	return root
}

func inOrder(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}

	result = inOrder(root.Left, result)
	result = append(result, root.Val)
	result = inOrder(root.Right, result)

	return result
}

func main() {
	nums := []int{9, 8, 2, 6, 1, -3, 7}
	fmt.Println(sort(nums))
}
