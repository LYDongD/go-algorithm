package main

import "fmt"
import "github.com/golang-collections/collections/stack"

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

	mergeStack := stack.New()
	nodePair := []*TreeNode{t1, t2}
	mergeStack.Push(nodePair)

	for mergeStack.Len() != 0 {
		nodePair = mergeStack.Pop().([]*TreeNode)
		if nodePair[0] == nil || nodePair[1] == nil {
			continue
		}
		nodePair[0].Val += nodePair[1].Val
		if nodePair[0].Left == nil {
			nodePair[0].Left = nodePair[1].Left
		} else {
			mergeStack.Push([]*TreeNode{nodePair[0].Left, nodePair[1].Left})
		}

		if nodePair[0].Right == nil {
			nodePair[0].Right = nodePair[1].Right
		} else {
			mergeStack.Push([]*TreeNode{nodePair[0].Right, nodePair[0].Right})
		}

	}

	return t1
}

func main() {
	fmt.Println("haha")
}
