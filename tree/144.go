package main

import "fmt"

type Stack struct {
    nodes []*TreeNode
}

func (self *Stack) isEmpty() bool {
    return len(self.nodes) == 0
}

func (self *Stack) push(node *TreeNode) {
    self.nodes = append(self.nodes, node)
}

func (self *Stack) pop() *TreeNode {
    if self.isEmpty() {
	panic("stack is empty")
    }

    top := self.nodes[len(self.nodes)-1]
    self.nodes = self.nodes[:len(self.nodes)-1]
    return top
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
    result := []int{}
    if root == nil {
	return result
    }

    stack := &Stack{nodes : []*TreeNode{}}
    stack.push(root)
    for !stack.isEmpty() {
	top := stack.pop()
	result = append(result, top.Val)
	if top.Right != nil {
	    stack.push(top.Right)
	}

	if top.Left != nil {
	    stack.push(top.Left)
	}
    }

    return result
}

func main() {
    fmt.Println("haha")
}


