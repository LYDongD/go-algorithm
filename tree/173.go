package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
	top := self.nodes[len(self.nodes)-1]
	self.nodes = self.nodes[:len(self.nodes)-1]
	return top
}

type BSTIterator struct {
	stack *Stack
}

func (this *BSTIterator) pushLeft(node *TreeNode) {
	for node != nil {
		this.stack.push(node)
		node = node.Left
	}
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := &BSTIterator{stack: &Stack{nodes: []*TreeNode{}}}
	iterator.pushLeft(root)
	return *iterator
}

func (this *BSTIterator) Next() int {
	if this.stack.isEmpty() {
		panic("no more nodes")
	}
	top := this.stack.pop()
	if top.Right != nil {
		this.pushLeft(top.Right)
	}
	return top.Val
}

func (this *BSTIterator) HasNext() bool {
	return !this.stack.isEmpty()
}

func main() {
	fmt.Println("vim-go")
}
