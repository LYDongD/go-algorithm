//package main
//
//import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {

	return getMinDiff(root, -1)
}

func getMinDiff(root *TreeNode, minDiff int) int {

	if root == nil {
		return minDiff
	}

	leftMin := getMinFromLeftTree(root.Left)
	rightMin := getMinFromRightTree(root.Right)

	if leftMin > 0 && (minDiff > root.Val-leftMin || minDiff < 0) {
		minDiff = root.Val - leftMin
	}

	if rightMin > 0 && (minDiff > rightMin-root.Val || minDiff < 0) {
		minDiff = rightMin - root.Val
	}

	minDiff = getMinDiff(root.Left, minDiff)
	minDiff = getMinDiff(root.Right, minDiff)

	return minDiff
}

func getMinFromLeftTree(root *TreeNode) int {
	if root == nil {
		return -1
	}

	if root.Right == nil {
		return root.Val
	}

	return getMinFromLeftTree(root.Right)
}

func getMinFromRightTree(root *TreeNode) int {
	if root == nil {
		return -1
	}

	if root.Left == nil {
		return root.Val
	}

	return getMinFromRightTree(root.Left)
}

//func main() {
//	root := &TreeNode{Val: 1}
//	root.Right = &TreeNode{Val: 3}
//	root.Right.Left = &TreeNode{Val: 2}
//	fmt.Println(getMinimumDifference(root))
//}
