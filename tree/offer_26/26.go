package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {

	if B == nil {
		return false
	}

	rootAList := findRootBInA(A, B.Val)
	for _, rootA := range rootAList {
		same := compare(rootA, B)
		if same {
			return true
		}
	}

	return false
}

func findRootBInA(A *TreeNode, rootBVal int) []*TreeNode {
	result := []*TreeNode{}
	return dfs(A, rootBVal, result)
}

func dfs(A *TreeNode, rootBVal int, result []*TreeNode) []*TreeNode {
	if A == nil {
		return result
	}

	if A.Val == rootBVal {
		result = append(result, A)
	}

	result = dfs(A.Left, rootBVal, result)
	result = dfs(A.Right, rootBVal, result)
	return result
}

func compare(rootA *TreeNode, rootB *TreeNode) bool {
	if rootB == nil {
		return true
	}

	if rootA == nil || rootA.Val != rootB.Val {
		return false
	}

	leftResult := compare(rootA.Left, rootB.Left)
	rightResult := compare(rootA.Right, rootB.Right)
	return leftResult && rightResult
}

func main() {
	A := &TreeNode{Val: 3}
	A.Left = &TreeNode{Val: 4}
	A.Right = &TreeNode{Val: 5}
	A.Left.Left = &TreeNode{Val: 1}
	A.Left.Right = &TreeNode{Val: 2}

	B := &TreeNode{Val: 3}
	//	B.Left = &TreeNode{Val: 4}
	B.Right = &TreeNode{Val: 5}

	fmt.Println(isSubStructure(A, B))
}
