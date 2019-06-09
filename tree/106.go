package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	inorderDic := make(map[int]int)
	for index, num := range inorder {
		inorderDic[num] = index
	}

	return buildTreeRecursively(inorderDic, 0, len(postorder)-1, len(postorder)-1, postorder)
}

func buildTreeRecursively(inorderDic map[int]int, start, end, current int, postorder []int) *TreeNode {

	if start > end || current < 0 {
		return nil
	}

	//get inorderindex from dic
	inorderIndex, ok := inorderDic[postorder[current]]
	if !ok {
		return nil
	}

	node := &TreeNode{Val: postorder[current]}
	node.Right = buildTreeRecursively(inorderDic, inorderIndex+1, end, current-1, postorder)
	node.Left = buildTreeRecursively(inorderDic, start, inorderIndex-1, current-(end-inorderIndex+1), postorder)
	return node
}

func main() {
	node := buildTree([]int{3, 2, 1}, []int{3, 2, 1})
	fmt.Println(node.Val, node.Left, node.Right)
	//fmt.Println(node.Right.Left, node.Right.Right.Val)
}
