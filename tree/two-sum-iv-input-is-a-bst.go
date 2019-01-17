package tree

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {

	goalMap := make(map[int]int)

	return checkTarget(root, goalMap, k)

}

func checkTarget(root *TreeNode, goalMap map[int]int, k int) bool {
	if root == nil {
		return false
	}

	if goalMap[root.Val] > 0 {
		return true
	}

	goalMap[k-root.Val]++

	return checkTarget(root.Left, goalMap, k) || checkTarget(root.Right, goalMap, k)

}

func main() {
	fmt.Println("vim-go")
}
