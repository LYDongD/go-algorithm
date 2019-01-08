package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	var result []int

	mergeTreesRecursively(t1, t2, result)

	fmt.Println("结果:", result)

	return buildTree(result)

}

func mergeTreesRecursively(t1 *TreeNode, t2 *TreeNode, result []int) {

	if t1 == nil && t2 == nil {
		return
	}

	if t1 == nil {
		result = append(result, t2.Val)
		mergeTreesRecursively(nil, t2.Left, result)
		mergeTreesRecursively(nil, t2.Right, result)
	} else if t2 == nil {
		result = append(result, t1.Val)
		mergeTreesRecursively(t1.Left, nil, result)
		mergeTreesRecursively(t2.Right, nil, result)
	} else {
		result = append(result, t1.Val+t2.Val)
		mergeTreesRecursively(t1.Left, t2.Left, result)
		mergeTreesRecursively(t1.Right, t2.Right, result)
	}

	fmt.Println(result)
}

func buildTree(result []int) *TreeNode {
	index := 0
	return buildTreePreorder(result, &index)
}

func buildTreePreorder(result []int, index *int) *TreeNode {
	if *index >= len(result) {
		return nil
	}

	root := &TreeNode{Val: result[*index]}

	*index = *index + 1
	root.Left = buildTreePreorder(result, index)
	root.Right = buildTreePreorder(result, index)

	return root

}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Print(root.Val)
	preorder(root.Left)
	preorder(root.Right)
}

func main() {

	tree1 := []int{1, 3, 5, 2}
	tree2 := []int{2, 1, 4, 3, 7}
	preorder(mergeTrees(buildTree(tree1), buildTree(tree2)))
}
