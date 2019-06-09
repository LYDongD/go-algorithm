package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	preOrderDic := make(map[int]int)
	for index, num := range preorder {
		preOrderDic[num] = index
	}

	inOrderDic := make(map[int]int)
	for index, num := range inorder {
		inOrderDic[num] = index
	}

	selectedDic := make(map[int]int)

	root := &TreeNode{Val: preorder[0]}
	selectedDic[preorder[0]]++

	generateTree(root, inOrderDic, preOrderDic, selectedDic, preorder, inorder)

	return root
}

func generateTree(node *TreeNode, inOrderDic, preOrderDic, selectedDic map[int]int, preorder, inorder []int) {

	index := inOrderDic[node.Val]
	selecteIndex, preOrderIndex := -1, -1
	//select left node
	for i := index - 1; i >= 0; i-- {
		if selectedDic[inorder[i]] > 0 {
			break
		}

		if preOrderIndex == -1 || preOrderDic[inorder[i]] < preOrderIndex {
			selecteIndex = i
			preOrderIndex = preOrderDic[inorder[i]]
		}
	}

	if selecteIndex >= 0 {
		node.Left = &TreeNode{Val: inorder[selecteIndex]}
		selectedDic[inorder[selecteIndex]]++
		generateTree(node.Left, inOrderDic, preOrderDic, selectedDic, preorder, inorder)
	}

	//select right node
	selecteIndex, preOrderIndex = -1, -1
	for i := index + 1; i < len(inorder); i++ {
		if selectedDic[inorder[i]] > 0 {
			break
		}

		if preOrderIndex == -1 || preOrderDic[inorder[i]] < preOrderIndex {
			selecteIndex = i
			preOrderIndex = preOrderDic[inorder[i]]
		}
	}

	if selecteIndex >= 0 {
		node.Right = &TreeNode{Val: inorder[selecteIndex]}
		selectedDic[inorder[selecteIndex]]++
		generateTree(node.Right, inOrderDic, preOrderDic, selectedDic, preorder, inorder)
	}
}

func main() {
	node := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(node.Val, node.Left.Val, node.Right.Val)
	fmt.Println(node.Right.Left.Val, node.Right.Right.Val)
}
