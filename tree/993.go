package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Trace struct {
	Val   int
	Level int
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil || root.Left == nil || root.Right == nil {
		return false
	}

	if root.Val == x || root.Val == y {
		return false
	}

	traceDic := make(map[int]*Trace)
	return preOrder(nil, root, x, y, 0, traceDic)
}

func preOrder(parent *TreeNode, node *TreeNode, x, y, level int, traceDic map[int]*Trace) bool {
	if node == nil {
		return false
	}

	var trace *Trace
	if node.Val == x {
		trace = traceDic[y]
	} else if node.Val == y {
		trace = traceDic[x]
	}

	if node.Val == x || node.Val == y {
		if trace == nil {
			traceDic[node.Val] = &Trace{parent.Val, level + 1}
		} else {
			if trace.Val != parent.Val && trace.Level == level+1 {
				return true
			} else {
				return false
			}
		}
	}

	return preOrder(node, node.Left, x, y, level+1, traceDic) || preOrder(node, node.Right, x, y, level+1, traceDic)

}

func main() {
	fmt.Println("vim-go")
}
