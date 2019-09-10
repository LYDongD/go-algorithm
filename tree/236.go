package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	pDic := findPath(root, p, 1, make(map[*TreeNode]int))
	qDic := findPath(root, q, 1, make(map[*TreeNode]int))
	return findAncestor(pDic, qDic)
}

func findPath(root, node *TreeNode, k int, path map[*TreeNode]int) map[*TreeNode]int {
	if root == node {
		path[root] = k
		return path
	}

	if root.Left != nil {
		tmpPath := copyMap(path)
		tmpPath[root] = k
		path = findPath(root.Left, node, k+1, tmpPath)
	}

	if root.Right != nil && path[node] == 0 {
		tmpPath := copyMap(path)
		tmpPath[root] = k
		path = findPath(root.Right, node, k+1, tmpPath)
	}

	return path
}

func findAncestor(pDic, qDic map[*TreeNode]int) *TreeNode {
	var ancestor *TreeNode
	depth := 0
	for key, value := range pDic {
		currentDepth := min(value, qDic[key])
		if qDic[key] > 0 && currentDepth > depth {
			ancestor = key
			depth = currentDepth
		}
	}

	return ancestor
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func copyMap(dic map[*TreeNode]int) map[*TreeNode]int {
	newDic := make(map[*TreeNode]int)
	for key, value := range dic {
		newDic[key] = value
	}

	return newDic
}

func main() {
	fmt.Println("haha")
}
