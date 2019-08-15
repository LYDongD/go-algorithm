package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LevelNode struct {
	node  *TreeNode
	level int
}

type Queue struct {
	nodes []*LevelNode
}

func (self *Queue) isEmpty() bool {
	return len(self.nodes) == 0
}

func (self *Queue) offer(node *LevelNode) {
	self.nodes = append(self.nodes, node)
}

func (self *Queue) take() *LevelNode {
	if self.isEmpty() {
		panic("queue is empty")
	}

	head := self.nodes[0]
	self.nodes = self.nodes[1:]
	return head
}

func rightSideView(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	queue := &Queue{nodes: []*LevelNode{}}
	queue.offer(&LevelNode{root, 1})
	currentLevel := 1
	var currentNode *TreeNode
	for !queue.isEmpty() {
		levelNode := queue.take()
		if levelNode.level != currentLevel {
			result = append(result, currentNode.Val)
		}

		currentLevel = levelNode.level
		currentNode = levelNode.node

		if currentNode.Left != nil {
			queue.offer(&LevelNode{currentNode.Left, levelNode.level + 1})
		}

		if currentNode.Right != nil {
			queue.offer(&LevelNode{currentNode.Right, levelNode.level + 1})
		}
	}

	//handle last level
	result = append(result, currentNode.Val)

	return result
}
func main() {

	fmt.Println("vim-go")
}
