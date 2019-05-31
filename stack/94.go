package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	nums []*TreeNode
}

func (self *Stack) isEmpty() bool {
	return len(self.nums) == 0
}

func (self *Stack) push(num *TreeNode) {
	self.nums = append(self.nums, num)
}

func (self *Stack) pop() *TreeNode {
	if self.isEmpty() {
		panic("stack is empty!!")
	}

	popNum := self.nums[len(self.nums)-1]
	self.nums = self.nums[:len(self.nums)-1]
	return popNum
}

func inorderTraversal(root *TreeNode) []int {

	result := []int{}

	if root == nil {
		return result
	}

	stack := &Stack{nums: []*TreeNode{}}

	current := root
	for current != nil || !stack.isEmpty() {
		//push left
		for current != nil {
			stack.push(current)
			current = current.Left
		}

		popEle := stack.pop()
		result = append(result, popEle.Val)

		//push right
		if popEle.Right != nil {
			current = popEle.Right
		}

	}

	return result
}

func main() {
	root := &TreeNode{1, nil, nil}
	root.Right = &TreeNode{2, nil, nil}
	root.Right.Left = &TreeNode{3, nil, nil}
	fmt.Println(inorderTraversal(root))
}
